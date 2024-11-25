package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"

	"golang.org/x/image/draw"
)

// Function to resize the image
func ResizeImage(b64Str string, newWidth int) string {
	// Decode the base64 string
	data, err := base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		panic(err)
	}

	// Decode the image
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	// Calculate new height to maintain aspect ratio
	newHeight := int(float64(newWidth) * float64(img.Bounds().Dy()) / float64(img.Bounds().Dx()))

	// Create a new image with the new dimensions
	resizedImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.BiLinear.Scale(resizedImg, resizedImg.Rect, img, img.Bounds(), draw.Over, nil)

	// Encode the resized image to PNG
	var buf bytes.Buffer
	if err := png.Encode(&buf, resizedImg); err != nil {
		panic(err)
	}

	// Return the base64 encoded string of the resized image
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
