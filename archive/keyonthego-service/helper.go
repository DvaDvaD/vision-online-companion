package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func GetIPAddress() ([]NetworkInterface, error) {
	if viper.GetString("MODE") == "cloud" {
		return []NetworkInterface{
			{
				Name: "cloud",
				IP:   "service.portierglobal.com",
			},
		}, nil
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Error().Err(err).Msg("Error getting interfaces")
		return nil, err
	}

	var validInterfaces []NetworkInterface

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 || strings.Contains(iface.Name, "Bluetooth") || strings.Contains(iface.Name, "WSL") {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			log.Error().Err(err).Str("interface", iface.Name).Msg("Error getting addresses")
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			ip := ipNet.IP
			if ip.To4() == nil {
				continue
			}

			validInterfaces = append(validInterfaces, NetworkInterface{
				Name: iface.Name,
				IP:   ip.String(),
			})
		}
	}

	return validInterfaces, nil
}

func GenerateULID() string {
	entropy := ulid.Monotonic(rand.Reader, 0)
	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}

func GenerateToken(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, length)
	for i := range token {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			log.Error().Err(err).Msg("Error generating random token")
			return ""
		}
		token[i] = charset[n.Int64()]
	}
	return string(token)
}

func PersistData(requestID, token string, body SignDataBody) error {
	// Create a hash of the token
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])

	// Create the content structure
	content := SignData{
		Token: tokenHash,
		Body:  body,
	}

	// Convert the content to JSON
	jsonData, err := json.Marshal(content)
	if err != nil {
		log.Error().Err(err).Msg("Error marshaling JSON")
		return err
	}

	// Ensure the directory exists
	err = os.MkdirAll(viper.GetString("TMP_FOLDER"), os.ModePerm)
	if err != nil {
		log.Error().Err(err).Msg("Error creating directory")
		return err
	}

	// Create the file
	filename := filepath.Join(viper.GetString("TMP_FOLDER"), requestID+".json")
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Error().Err(err).Msg("Error writing file")
		return err
	}

	log.Info().Str("filename", filename).Msg("Request file created")
	return nil
}

func VerifyTokenAndGetContent(requestID, token string) (*SignData, error) {
	// Read the file
	data, err := os.ReadFile(filepath.Join(viper.GetString("TMP_FOLDER"), requestID+".json"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("request id or token invalid")
		}
		return nil, fmt.Errorf("error reading sign request")
	}

	// Unmarshal the JSON data
	var content SignData
	if err := json.Unmarshal(data, &content); err != nil {
		return nil, fmt.Errorf("error parsing sign request")
	}

	// Verify the token
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])
	if tokenHash != content.Token {
		return nil, fmt.Errorf("request id or token invalid")
	}

	return &content, nil
}
