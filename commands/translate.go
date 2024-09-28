package commands

import (
	"github.com/LordPax/aicli/lang"
	"github.com/LordPax/aicli/sdk"
	"github.com/LordPax/aicli/service"
	cli "github.com/urfave/cli/v2"
)

func TranslateCommand() (*cli.Command, error) {
	l := lang.GetLocalize()

	if err := sdk.InitSdkTranslate(""); err != nil {
		return nil, err
	}

	return &cli.Command{
		Name:      "translate",
		Usage:     l.Get("translate-usage"),
		ArgsUsage: "[text|-]",
		Aliases:   []string{"tr"},
		Action:    translateAction,
		Flags:     translateFlags(),
	}, nil
}

func translateFlags() []cli.Flag {
	l := lang.GetLocalize()
	sdkTranslate := sdk.GetSdkTranslate()

	return []cli.Flag{
		&cli.StringFlag{
			Name:        "sdk",
			Aliases:     []string{"S"},
			Usage:       l.Get("sdk-usage"),
			DefaultText: sdkTranslate.GetName(),
			Category:    "global",
			Action: func(c *cli.Context, value string) error {
				if err := sdk.InitSdkTranslate(value); err != nil {
					return err
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:     "source",
			Aliases:  []string{"s"},
			Usage:    l.Get("translate-source-usage"),
			Category: "translate",
			Action: func(c *cli.Context, value string) error {
				sdkTranslate := sdk.GetSdkTranslate()
				sdkTranslate.SetSourceLang(value)
				return nil
			},
		},
		&cli.StringFlag{
			Name:     "target",
			Aliases:  []string{"t"},
			Usage:    l.Get("translate-target-usage"),
			Category: "translate",
			Action: func(c *cli.Context, value string) error {
				sdkTranslate := sdk.GetSdkTranslate()
				sdkTranslate.SetTargetLang(value)
				return nil
			},
		},
	}
}

func translateAction(c *cli.Context) error {
	text := c.Args().First()

	if c.NArg() == 0 {
		if err := service.TranslateInteractiveMode(); err != nil {
			return err
		}
		return nil
	}

	return service.TranslateText(text)
}
