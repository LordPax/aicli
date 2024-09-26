package commands

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/LordPax/aicli/lang"
	"github.com/LordPax/aicli/sdk"
	"github.com/LordPax/aicli/service"

	cli "github.com/urfave/cli/v2"
)

func TextCommand() *cli.Command {
	l := lang.GetLocalize()
	return &cli.Command{
		Name:      "text",
		Usage:     l.Get("text-usage"),
		ArgsUsage: "[prompt|-]",
		Aliases:   []string{"t"},
		Action:    textAction,
		Flags:     textFlags(),
	}
}

func textFlags() []cli.Flag {
	l := lang.GetLocalize()
	textSdk := sdk.GetSdkText()

	return []cli.Flag{
		&cli.StringFlag{
			Name:        "sdk",
			Aliases:     []string{"S"},
			Usage:       l.Get("sdk-usage"),
			DefaultText: textSdk.GetName(),
			Action: func(c *cli.Context, value string) error {
				if err := sdk.InitSdkText(value); err != nil {
					return err
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "history",
			Aliases:     []string{"H"},
			Usage:       l.Get("text-history-usage"),
			DefaultText: textSdk.GetSelectedHistory(),
			Action: func(c *cli.Context, value string) error {
				text := sdk.GetSdkText()
				text.SetSelectedHistory(value)
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "model",
			Aliases:     []string{"m"},
			Usage:       l.Get("sdk-model-usage"),
			DefaultText: textSdk.GetModel(),
			Action: func(c *cli.Context, value string) error {
				text := sdk.GetSdkText()
				text.SetModel(value)
				return nil
			},
		},
		&cli.Float64Flag{
			Name:        "temp",
			Aliases:     []string{"t"},
			Usage:       l.Get("text-temp-usage"),
			DefaultText: strconv.FormatFloat(textSdk.GetTemp(), 'f', -1, 64),
			Action: func(c *cli.Context, value float64) error {
				text := sdk.GetSdkText()
				text.SetTemp(value)
				return nil
			},
		},
		&cli.StringSliceFlag{
			Name:    "system",
			Aliases: []string{"s"},
			Usage:   l.Get("text-system-usage"),
			Action: func(c *cli.Context, values []string) error {
				text := sdk.GetSdkText()
				var content []string

				for _, value := range values {
					if value == "-" {
						stdin, err := io.ReadAll(os.Stdin)
						if err != nil {
							return err
						}

						value = string(stdin)
					}

					content = append(content, value)
				}

				text.AppendHistory("system", content...)

				return nil
			},
		},
		&cli.StringSliceFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   l.Get("text-file-usage"),
			Action: func(c *cli.Context, files []string) error {
				text := sdk.GetSdkText()
				var fileContent []string

				for _, file := range files {
					f, err := os.ReadFile(file)
					if err != nil {
						return err
					}

					if len(f) == 0 {
						return fmt.Errorf(l.Get("empty-file"), file)
					}

					fileContent = append(fileContent, string(f))
				}

				text.AppendHistory("system", fileContent...)

				return nil
			},
		},
		&cli.BoolFlag{
			Name:               "clear",
			Aliases:            []string{"c"},
			Usage:              l.Get("text-clear-usage"),
			DisableDefaultText: true,
			Action: func(c *cli.Context, value bool) error {
				text := sdk.GetSdkText()
				text.ClearHistory()
				if err := text.SaveHistory(); err != nil {
					return err
				}
				os.Exit(0)
				return nil
			},
		},
		&cli.BoolFlag{
			Name:               "list-history",
			Aliases:            []string{"l"},
			Usage:              l.Get("text-list-history-usage"),
			DisableDefaultText: true,
			Action: func(c *cli.Context, value bool) error {
				if err := service.ListHistory(true); err != nil {
					return err
				}
				os.Exit(0)
				return nil
			},
		},
		&cli.BoolFlag{
			Name:               "list-history-name",
			Aliases:            []string{"L"},
			Usage:              l.Get("text-list-history-name-usage"),
			DisableDefaultText: true,
			Action: func(c *cli.Context, value bool) error {
				text := sdk.GetSdkText()
				for _, name := range text.GetHistoryNames() {
					fmt.Println(name)
				}
				os.Exit(0)
				return nil
			},
		},
	}
}

func textAction(c *cli.Context) error {
	if c.NArg() == 0 {
		if err := service.InteractiveMode(); err != nil {
			return err
		}
		return nil
	}

	if err := service.SendTextRequest(c.Args().First()); err != nil {
		return err
	}

	return nil
}
