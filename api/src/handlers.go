package main

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/portierglobal/vision-online-companion/api/gen"
	"github.com/portierglobal/vision-online-companion/business/keyonthego"
	"github.com/portierglobal/vision-online-companion/database/data"
)

func (app *application) GetKeyOtgSign(ctx echo.Context) error {
	signRequests, err := app.queries.GetAllSignRequests(context.Background())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	var genSignRequests []gen.SignResponse
	for _, signRequest := range signRequests {
		genSignRequests = append(genSignRequests, gen.SignResponse{
			RequestId:   &signRequest.RequestID,
			RequestUser: signRequest.RequestUser,
			HolderId:    signRequest.HolderID,
			HolderName:  signRequest.HolderName,
			Notes:       &signRequest.Notes.String,
			Status:      (*gen.Status)(&signRequest.Status.String),
			CreatedAt:   &signRequest.CreatedAt.Time,
			UpdatedAt:   &signRequest.UpdatedAt.Time,
			SignedAt:    &signRequest.SignedAt.Time,

			// TODO: Create join query to get issue, sign, and location fields
			LocationLatitude:  nil, // Assuming these fields are not available in signRequest
			LocationLongitude: nil, // Assuming these fields are not available in signRequest
			Sign:              nil, // Assuming these fields are not available in signRequest
			Issue:             nil, // Assuming these fields are not available in signRequest
		})
	}

	return ctx.JSON(http.StatusOK, genSignRequests)
}

func (app *application) PostKeyOtgSign(ctx echo.Context, params gen.PostKeyOtgSignParams) error {
	var req gen.CreateSignRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// Create sign request
	signRequestParams := data.CreateSignRequestParams{
		RequestUser: req.RequestUser,
		HolderID:    req.HolderId,
		HolderName:  req.HolderName,
		Notes:       pgtype.Text{String: *req.Notes, Valid: req.Notes != nil},
		Status:      pgtype.Text{String: "pending", Valid: true},
		Token:       app.generateToken(), // Assume generateToken() is a function that generates a token
	}
	signRequest, err := app.queries.CreateSignRequest(context.Background(), signRequestParams)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	// Create issues
	for _, issue := range req.Issue {
		issueParams := data.CreateIssueParams{
			SignRequestID: signRequest.RequestID,
			Number:        issue.Number,
			Copy:          int32(issue.Copy),
			Description:   pgtype.Text{String: *issue.Description, Valid: issue.Description != nil},
		}
		_, err := app.queries.CreateIssue(context.Background(), issueParams)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}

	response := gen.CreateSignResponse{
		RequestId: &signRequest.RequestID,
		Token:     &signRequest.Token,
	}
	return ctx.JSON(http.StatusCreated, response)
}

func (app *application) GetKeyOtgSignRequestID(ctx echo.Context, requestID gen.RequestID, params gen.GetKeyOtgSignRequestIDParams) error {
	signRequest, err := app.queries.GetSignRequestByID(context.Background(), requestID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	issues, err := app.queries.GetIssuesBySignRequestID(context.Background(), requestID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	var genIssues []gen.Issue
	for _, issue := range issues {
		genIssue := app.convertDataIssueToGenIssue(issue)
		genIssues = append(genIssues, genIssue)
	}

	response := gen.SignResponse{
		RequestId:   &signRequest.RequestID,
		RequestUser: signRequest.RequestUser,
		HolderId:    signRequest.HolderID,
		HolderName:  signRequest.HolderName,
		Notes:       &signRequest.Notes.String,
		Status:      (*gen.Status)(&signRequest.Status.String),
		Issue:       genIssues,
		CreatedAt:   &signRequest.CreatedAt.Time,
		UpdatedAt:   &signRequest.UpdatedAt.Time,
		SignedAt:    &signRequest.SignedAt.Time,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (app *application) PostKeyOtgSignRequestID(ctx echo.Context, requestID gen.RequestID, params gen.PostKeyOtgSignRequestIDParams) error {
	var req gen.SignSubmitRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	signSubmissionParams := data.CreateSignSubmissionParams{
		SignRequestID:     requestID,
		Sign:              []byte(req.Sign),
		LocationLatitude:  pgtype.Float8{Float64: float64(*req.LocationLatitude), Valid: req.LocationLatitude != nil},
		LocationLongitude: pgtype.Float8{Float64: float64(*req.LocationLongitude), Valid: req.LocationLongitude != nil},
	}
	_, err := app.queries.CreateSignSubmission(context.Background(), signSubmissionParams)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	updateStatusParams := data.UpdateSignRequestStatusParams{
		Status:    pgtype.Text{String: "signed", Valid: true},
		RequestID: requestID,
	}
	err = app.queries.UpdateSignRequestStatus(context.Background(), updateStatusParams)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (app *application) GetKeyOtgSignRequestIDQr(ctx echo.Context, requestID gen.RequestID, params gen.GetKeyOtgSignRequestIDQrParams) error {
	hostname := ctx.Scheme() + "://" + ctx.Request().Host
	token := params.Token

	png, err := keyonthego.CreateQR(hostname, token, string(requestID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	// Return the QR code image as a PNG
	return ctx.Blob(http.StatusOK, "image/png", png)
}

func (app *application) PostShutdown(ctx echo.Context) error {
	go func() {
		// Delay a bit before shutting down so the response can be sent
		time.Sleep(1 * time.Second)
		app.shutdownCancel() // Trigger server shutdown
	}()

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Server shutting down...",
	})
}
