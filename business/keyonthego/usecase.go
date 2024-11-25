package keyonthego

import (
	"fmt"
	"net/url"

	"github.com/skip2/go-qrcode"
)

func CreateQR(hostname, token, requestID string) ([]byte, error) {
	// Construct the full URL: http://hostname/sign/requestID?token=abcdef
	fullURL := fmt.Sprintf("%s/key-otg/sign/%s?token=%s", hostname, requestID, url.QueryEscape(token))

	// Generate the QR code for the full URL
	var png []byte
	png, err := qrcode.Encode(fullURL, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return png, nil
}
