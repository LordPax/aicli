package service

import (
	"io"
	"os"

	"github.com/LordPax/aicli/sdk"
)

func SendImageRequest(prompt string) error {
	imageSdk := sdk.GetSdkImage()

	if prompt == "-" {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		prompt = string(stdin)
	}

	if err := imageSdk.SendRequest(prompt); err != nil {
		return err
	}

	return nil
}
