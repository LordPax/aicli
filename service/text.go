package service

import (
	"fmt"

	"github.com/LordPax/aicli/sdk"
)

func SendTextRequest(prompt string) error {
	textSdk := sdk.GetSdkText()

	resp, err := textSdk.SendRequest(prompt)
	if err != nil {
		return err
	}

	fmt.Println(resp.Content)

	return nil
}
