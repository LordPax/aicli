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
	log, err := utils.GetLog()
	if err != nil {
		return err
	}

	if translateSdk.GetTargetLang() == "" {
		return errors.New(l.Get("translate-target-required"))
	}

	for {
		input := utils.Input(l.Get("translate-input"), "", false)
		ok, err := ParseTranslateCommand(input)
		if err != nil {
			log.PrintfErr("%v\n", err)
			continue
		}
		if ok {
			continue
		}

		resp, err := translateSdk.SendRequest(input)
		if err != nil {
			return err
		}

		fmt.Print("\n")
		fmt.Println(utils.Red + "> " + utils.Reset + resp)
		fmt.Print("\n")
	}
}

func ParseTranslateCommand(command string) (bool, error) {
	l := lang.GetLocalize()

	if command[0] != '/' {
		return false, nil
	}

	values := strings.Split(command, " ")

	switch values[0][1:] {
	case "help":
		HelpTranslateCommand()
		return true, nil
	case "exit", "e":
		os.Exit(0)
		return true, nil
	case "sdk", "s":
		translateSdk := sdk.GetSdkTranslate()
		if len(values) < 2 {
			fmt.Println(translateSdk.GetName())
			return true, nil
		}

		if err := sdk.InitSdkTranslate(values[1]); err != nil {
			return true, err
		}

		return true, nil
	case "source", "so":
		translateSdk := sdk.GetSdkTranslate()
		if len(values) < 2 {
			if translateSdk.GetSourceLang() == "" {
				fmt.Println(l.Get("translate-detect"))
				return true, nil
			}

			fmt.Println(translateSdk.GetSourceLang())
			return true, nil
		}

		if values[1] == "auto" {
			translateSdk.SetSourceLang("")
			fmt.Println(l.Get("translate-detect"))
			return true, nil
		}

		translateSdk.SetSourceLang(values[1])
		fmt.Printf(l.Get("translate-source-set"), values[1])
		return true, nil
	case "target", "ta":
		translateSdk := sdk.GetSdkTranslate()
		if len(values) < 2 {
			fmt.Println(translateSdk.GetTargetLang())
			return true, nil
		}

		translateSdk.SetTargetLang(values[1])
		fmt.Printf(l.Get("translate-target-set"), values[1])
		return true, nil
	}

	return false, fmt.Errorf(l.Get("unknown-command"), command)
}

func HelpTranslateCommand() {
	l := lang.GetLocalize()
	fmt.Println(l.Get("translate-command-help-usage"))
	fmt.Println(l.Get("translate-command-exit-usage"))
	fmt.Println(l.Get("translate-command-sdk-usage"))
	fmt.Println(l.Get("translate-command-source-usage"))
	fmt.Println(l.Get("translate-command-target-usage"))
}
