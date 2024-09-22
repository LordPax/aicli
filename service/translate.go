package service

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/LordPax/aicli/lang"
	"github.com/LordPax/aicli/sdk"
	"github.com/LordPax/aicli/utils"
)

func TranslateText(text string) error {
	sdkTranslate := sdk.GetSdkTranslate()
	l := lang.GetLocalize()

	if text == "-" {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		text = strings.TrimSpace(string(stdin))
	}

	if sdkTranslate.GetTargetLang() == "" {
		return errors.New(l.Get("translate-target-required"))
	}

	resp, err := sdkTranslate.SendRequest(text)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func TranslateInteractiveMode() error {
	translateSdk := sdk.GetSdkTranslate()
	l := lang.GetLocalize()

	if translateSdk.GetTargetLang() == "" {
		return errors.New(l.Get("translate-target-required"))
	}

	for {
		input := utils.Input(l.Get("translate-input"), "", false)
		if input == "exit" {
			break
		}

		resp, err := translateSdk.SendRequest(input)
		if err != nil {
			return err
		}

		fmt.Print("\n")
		fmt.Println(utils.Red + "> " + utils.Reset + resp)
		fmt.Print("\n")
	}

	return nil
}
