package commands

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/LordPax/aicli/lang"
	"github.com/LordPax/aicli/sdk"
	"github.com/LordPax/aicli/service"
	"github.com/LordPax/aicli/utils"

	cli "github.com/urfave/cli/v2"
)

func TextCommand() (*cli.Command, error) {
	l := lang.GetLocalize()

	if err := sdk.InitSdkText(""); err != nil {
		return nil, err
	}

	return &cli.Command{
		Name:      "text",
		Usage:     l.Get("text-usage"),
		ArgsUsage: "[prompt|-]",
		Aliases:   []string{"t"},
		Action:    textAction,
		Flags:     textFlags(),
	}, nil
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
			Category:    "global",
			Action: func(c *cli.Context, value string) error {
				if err := sdk.InitSdkText(value); err != nil {
					return err
				}
				return nil
			},
		},
		&cli.BoolFlag{
			Name:               "inerte",
			Aliases:            []string{"i"},
			Usage:              l.Get("inerte-usage"),
			DisableDefaultText: true,
			Category:           "global",
			Action: func(c *cli.Context, value bool) error {
				text := sdk.GetSdkText()
				text.SetInerte(value)
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "history",
			Aliases:     []string{"H"},
			Usage:       l.Get("text-history-usage"),
			DefaultText: textSdk.GetSelectedHistory(),
			Category:    "history",
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
			Category:    "text",
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
			Category:    "text",
			Action: func(c *cli.Context, value float64) error {
				text := sdk.GetSdkText()
				text.SetTemp(value)
				return nil
			},
		},
		&cli.StringFlag{
			Name:     "context",
			Aliases:  []string{"s"},
			Usage:    l.Get("text-system-usage"),
			Category: "text",
			Action: func(c *cli.Context, value string) error {
				text := sdk.GetSdkText()

				if value == "-" {
					stdin, err := io.ReadAll(os.Stdin)
					if err != nil {
						return err
					}

					value = string(stdin)
				}

				text.AppendHistory("system", value)

				return nil
			},
		},
		&cli.StringSliceFlag{
			Name:     "file",
			Aliases:  []string{"f"},
			Usage:    l.Get("text-file-usage"),
			Category: "text",
			Action: func(c *cli.Context, files []string) error {
				text := sdk.GetSdkText()

				for _, file := range files {
					f, err := os.ReadFile(file)
					if err != nil {
						return err
					}

					if len(f) == 0 {
						return fmt.Errorf(l.Get("empty-file"), file)
					}

					if fileType := utils.IsFileType(f, utils.IMAGE); fileType != "" {
						if err := text.AppendImageHistory("system", "image/"+fileType, f); err != nil {
							return err
						}
						continue
					}

					text.AppendHistory("system", string(f))
				}

				return nil
			},
		},
		&cli.StringSliceFlag{
			Name:     "url",
			Aliases:  []string{"u"},
			Usage:    l.Get("text-url-usage"),
			Category: "text",
			Action: func(c *cli.Context, urls []string) error {
				text := sdk.GetSdkText()

				for _, url := range urls {
					f, err := utils.GetFileFromUrl(url)
					if err != nil {
						return err
					}

					if len(f) == 0 {
						return fmt.Errorf(l.Get("empty-file"), url)
					}

					text.AppendHistory("system", string(f))
				}

				return nil
			},
		},
		&cli.BoolFlag{
			Name:               "clear",
			Aliases:            []string{"c"},
			Usage:              l.Get("text-clear-usage"),
			DisableDefaultText: true,
			Category:           "history",
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
			Category:           "history",
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
			Category:           "history",
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
	textSdk := sdk.GetSdkText()

	if c.NArg() == 0 && !textSdk.GetInerte() {
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
