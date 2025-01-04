package service

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

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
	log, err := utils.GetLog()
	if err != nil {
		return err
	}

	if err := ListHistory(false); err != nil {
		return err
	}

	for {
		input := utils.Input(l.Get("text-input"), "", false)
		ok, err := ParseTextCommand(input)
		if err != nil {
			log.PrintfErr("%v\n", err)
			continue
		}
		if ok {
			continue
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
}

func ParseTextCommand(command string) (bool, error) {
	l := lang.GetLocalize()

	if command[0] != '/' {
		return false, nil
	}

	values := strings.Split(command, " ")

	switch values[0][1:] {
	case "help":
		HelpTextCommand()
		return true, nil
	case "exit", "e":
		os.Exit(0)
		return true, nil
	case "sdk", "s":
		textSdk := sdk.GetSdkText()
		if len(values) < 2 {
			fmt.Println(textSdk.GetName())
			return true, nil
		}

		if err := sdk.InitSdkText(values[1]); err != nil {
			return true, err
		}

		return true, nil
	case "model", "m":
		textSdk := sdk.GetSdkText()
		if len(values) < 2 {
			fmt.Println(textSdk.GetModel())
			return true, nil
		}

		textSdk.SetModel(values[1])

		return true, nil
	case "temp", "t":
		textSdk := sdk.GetSdkText()
		if len(values) < 2 {
			fmt.Println(textSdk.GetTemp())
			return true, nil
		}

		temp, err := strconv.ParseFloat(values[1], 64)
		if err != nil {
			return true, err
		}

		textSdk.SetTemp(temp)
		return true, nil
	case "history", "h":
		textSdk := sdk.GetSdkText()
		if len(values) < 2 {
			textSdk.ListHistoryNames()
			return true, nil
		}

		textSdk.SetSelectedHistory(values[1])
		fmt.Printf(l.Get("text-history-selected"), values[1])
		return true, nil
	case "show":
		if err := ListHistory(true); err != nil {
			return true, err
		}

		return true, nil
	case "clear", "c":
		textSdk := sdk.GetSdkText()
		textSdk.ClearHistory()
		if err := textSdk.SaveHistory(); err != nil {
			return true, err
		}

		fmt.Printf(l.Get("text-history-cleared"), textSdk.GetSelectedHistory())
		return true, nil
	}

	return false, fmt.Errorf(l.Get("unknown-command"), command)
}

func HelpTextCommand() {
	l := lang.GetLocalize()
	fmt.Println(l.Get("text-command-help-usage"))
	fmt.Println(l.Get("text-command-exit-usage"))
	fmt.Println(l.Get("text-command-sdk-usage"))
	fmt.Println(l.Get("text-command-model-usage"))
	fmt.Println(l.Get("text-command-temp-usage"))
	fmt.Println(l.Get("text-command-history-usage"))
	fmt.Println(l.Get("text-command-clear-usage"))
	fmt.Println(l.Get("text-command-show-usage"))
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
