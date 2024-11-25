//go:build !windows
// +build !windows

package utils

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

func GetTMPPath() string {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	fmt.Println(dirname)
	return dirname
}

// SendNotification does nothing on non-Windows systems
func SendNotification(holderName, requestID string) error {
	// No-op for non-Windows systems
	return nil
}
