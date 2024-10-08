package service

import (
	"fmt"
	"io"
	"os"

	"github.com/LordPax/aicli/lang"
	"github.com/LordPax/aicli/sdk"
	"github.com/LordPax/aicli/utils"
)

func SendTextRequest(prompt string) error {
	textSdk := sdk.GetSdkText()

	if prompt == "-" {
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		prompt = string(stdin)
	}

	resp, err := textSdk.SendRequest(prompt)
	if err != nil {
		return err
	}

	if resp.IsEmpty() {
		return nil
	}

	fmt.Println(resp.GetContent())

	return nil
}

func InteractiveMode() error {
	textSdk := sdk.GetSdkText()
	l := lang.GetLocalize()

	if err := ListHistory(false); err != nil {
		return err
	}

	for {
		input := utils.Input(l.Get("text-input"), "", false)
		if input == "exit" {
			break
		}

		resp, err := textSdk.SendRequest(input)
		if err != nil {
			return err
		}

		if resp.IsEmpty() {
			continue
		}

		fmt.Print("\n")
		fmt.Println(utils.Red + resp.Role + ">" + utils.Reset)
		fmt.Println(resp.GetContent())
		fmt.Print("\n")
	}

	return nil
}

func ListHistory(showMsg bool) error {
	textSdk := sdk.GetSdkText()
	l := lang.GetLocalize()
	log, err := utils.GetLog()
	if err != nil {
		return err
	}

	history := textSdk.GetHistory()

	if len(history) == 0 && showMsg {
		log.Printf(l.Get("empty-history"), textSdk.GetSelectedHistory())
		return nil
	}

	for _, message := range history {
		role := message.Role

		switch role {
		case "user":
			fmt.Print(utils.Blue + "user> " + utils.Reset)
		case "system":
			fmt.Println(utils.Green + "system> " + utils.Reset)
		case "assistant":
			fmt.Println(utils.Red + "assistant> " + utils.Reset)
		}

		fmt.Println(message.GetContent())
		fmt.Print("\n")
	}

	return nil
}
