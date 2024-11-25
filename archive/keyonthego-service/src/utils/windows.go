//go:build windows
// +build windows

package utils

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	"gopkg.in/toast.v1"
)

func GetTMPPath() string {
	dirname, err := os.UserConfigDir()

	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	dirname += "\\portierVision\\5"
	fmt.Println(dirname)
	return dirname
}

// SendNotification sends a notification for the signed request
func SendNotification(holderName, requestID string) error {
	notification := toast.Notification{
		AppID:   "portier Vision - KeyOnTheGo",
		Title:   fmt.Sprintf(`Request has been signed by %s`, holderName),
		Message: fmt.Sprintf(`Please check the request status from portier Vision with id: %s`, requestID),
		// Icon:    "go.png", // This file must exist (remove this line if it doesn't)
		// Actions: []toast.Action{
		// 	{"protocol", "I'm a button", ""},
		// 	{"protocol", "Me too!", ""},
		// },
	}
	return notification.Push()
}
