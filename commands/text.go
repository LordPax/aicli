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
		Flags:     TextFlags(),
	}
}

func TextFlags() []cli.Flag {
	l := lang.GetLocalize()
	textSdk := sdk.GetSdkText()

	return []cli.Flag{
		&cli.StringFlag{
			Name:        "history",
			Aliases:     []string{"H"},
			Usage:       l.Get("text-history-usage"),
			DefaultText: textSdk.GetSelectedHistory(),
			Action: func(c *cli.Context, value string) error {
				textSdk.SetSelectedHistory(value)
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "model",
			Aliases:     []string{"m"},
			Usage:       l.Get("sdk-model-usage"),
			DefaultText: textSdk.GetModel(),
			Action: func(c *cli.Context, value string) error {
				textSdk.SetModel(value)
				return nil
			},
		},
		&cli.Float64Flag{
			Name:        "temp",
			Aliases:     []string{"t"},
			Usage:       l.Get("text-temp-usage"),
			DefaultText: strconv.FormatFloat(textSdk.GetTemp(), 'f', -1, 64),
			Action: func(c *cli.Context, value float64) error {
				textSdk.SetTemp(value)
				return nil
			},
		},
		&cli.StringSliceFlag{
			Name:    "system",
			Aliases: []string{"s"},
			Usage:   l.Get("text-system-usage"),
			Action: func(c *cli.Context, values []string) error {
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

				textSdk.AppendHistory("system", content...)

				return nil
			},
		},
		&cli.StringSliceFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   l.Get("text-file-usage"),
			Action: func(c *cli.Context, files []string) error {
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

				textSdk.AppendHistory("system", fileContent...)

				return nil
			},
		},
		&cli.BoolFlag{
			Name:    "clear",
			Aliases: []string{"c"},
			Usage:   l.Get("text-clear-usage"),
			Action: func(c *cli.Context, value bool) error {
				textSdk.ClearHistory()
				if err := textSdk.SaveHistory(); err != nil {
					return err
				}
				os.Exit(0)
				return nil
			},
		},
		&cli.BoolFlag{
			Name:    "list-history",
			Aliases: []string{"l"},
			Usage:   l.Get("text-list-history-usage"),
			Action: func(c *cli.Context, value bool) error {
				if err := service.ListHistory(true, true); err != nil {
					return err
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
