package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/portierglobal/keyonthego-service/src/utils"
	"github.com/rs/zerolog/log"
	"github.com/skip2/go-qrcode"
	"github.com/spf13/viper"
)

// Handler for POST /sign
func CreateSign(c echo.Context) error {
	signRequest := new(CreateSignRequest)
	if err := c.Bind(signRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(signRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	networks, err := GetIPAddress()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	requestID := GenerateULID()
	token := GenerateToken(16)
	err = PersistData(requestID, token, SignDataBody{
		CreateSignRequest: *signRequest,
		Sign:              "",
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Modify the networks slice to include the formatted string
	dsnFormat := "http://%s:%d/key-otg/sign/%s?token=%s"
	if viper.GetString("MODE") == "cloud" {
		dsnFormat = "https://%s/api/v1/key-otg/sign/%s?token=%s"
		network := networks[0]
		networks[0].Dsn = fmt.Sprintf(dsnFormat, network.IP, requestID, token)
	} else {
		for i, network := range networks {
			networks[i].Dsn = fmt.Sprintf(dsnFormat, network.IP, Port, requestID, token)
		}
	}

	response := CreateSignResponse{
		RequestID:  requestID,
		Token:      token,
		Interfaces: networks,
	}

	return c.JSON(http.StatusOK, response)
}

// Handler for POST /sign/:requestID
func SubmitSign(c echo.Context) error {
	log.Info().Msgf("SubmitSign: %+v", c.Request().Header)
	requestID := c.Param("requestID")
	token := c.QueryParam("token")
	signSubmit := new(SignSubmitRequest)
	if err := c.Bind(signSubmit); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(signSubmit); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	content, err := VerifyTokenAndGetContent(requestID, token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	// Update the sign data
	content.Body.Sign = utils.ResizeImage(signSubmit.Sign, 600)

	// Marshal the updated content
	updatedData, err := json.Marshal(content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error updating sign request")
	}

	// Write the updated data back to the file
	filePath := filepath.Join(viper.GetString("TMP_FOLDER"), requestID+".json")
	if err := os.WriteFile(filePath, updatedData, 0644); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error saving sign request")
	}

	// Prepare the response
	response := SignResponse{
		CreateSignRequest: CreateSignRequest{
			RequestUser: content.Body.RequestUser,
			HolderID:    content.Body.HolderID,
			HolderName:  content.Body.HolderName,
			Notes:       content.Body.Notes,
			Issue:       content.Body.Issue,
		},
		SignSubmitRequest: *signSubmit,
		RequestID:         requestID,
		Status:            StatusSuccess,
	}

	if err := utils.SendNotification(content.Body.HolderName, requestID); err != nil {
		log.Error().Msgf("Error sending notification: %v", err)
	}

	return c.JSON(http.StatusOK, response)
}

// Handler for GET /sign/:requestID
func GetSign(c echo.Context) error {
	requestID := c.Param("requestID")
	token := c.QueryParam("token")

	if requestID == "" || token == "" {
		return c.JSON(http.StatusBadRequest, "Missing requestID or token")
	}

	content, err := VerifyTokenAndGetContent(requestID, token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	bodyMap := content.Body
	// log.Info().Msgf("bodyMap: %+v", bodyMap)

	// Prepare the response
	response := SignResponse{
		CreateSignRequest: CreateSignRequest{
			RequestUser: bodyMap.RequestUser,
			HolderID:    bodyMap.HolderID,
			HolderName:  bodyMap.HolderName,
			Notes:       bodyMap.Notes,
			Issue:       bodyMap.Issue,
		},
		SignSubmitRequest: SignSubmitRequest{
			Sign: bodyMap.Sign,
		},
		RequestID: requestID,
		Status: func() Status {
			if bodyMap.Sign == "" {
				return StatusPending
			}
			return StatusSuccess
		}(),
	}

	return c.JSON(http.StatusOK, response)
}

// Handler for GET /sign/:requestID
func GetURLasQR(c echo.Context) error {
	// Get the hostname from the request
	hostname := c.Scheme() + "://" + c.Request().Host

	// Get the path parameter (requestID) and query parameter (token)
	requestID := c.Param("requestID")
	token := c.QueryParam("token")
	if requestID == "" || token == "" {
		return c.JSON(http.StatusBadRequest, "Missing requestID or token")
	}

	// Construct the full URL: http://hostname/sign/requestID?token=abcdef
	fullURL := fmt.Sprintf("%s/key-otg/sign/%s?token=%s", hostname, requestID, url.QueryEscape(token))
	if viper.GetString("MODE") == "cloud" {
		fullURL = fmt.Sprintf("https://service.portierglobal.com/api/v1/key-otg/sign/%s?token=%s", requestID, url.QueryEscape(token))
	}

	// Generate the QR code for the full URL
	var png []byte
	png, err := qrcode.Encode(fullURL, qrcode.Medium, 256)
	if err != nil {
		return err
	}

	// Return the QR code image as a PNG
	return c.Blob(http.StatusOK, "image/png", png)
}
