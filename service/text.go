package service

import (
	"github.com/LordPax/aicli/sdk"
)

func SendTextRequest(prompt string) error {
	textSdk := sdk.GetSdkText()

	_, err := textSdk.SendRequest(prompt)
	if err != nil {
		return err
	}

	return nil
}
