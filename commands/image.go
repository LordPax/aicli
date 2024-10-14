package commands

import (
	"strconv"

	"github.com/LordPax/aicli/lang"
	"github.com/LordPax/aicli/sdk"
	cli "github.com/urfave/cli/v2"
)

func ImageCommand() (*cli.Command, error) {
	l := lang.GetLocalize()

	if err := sdk.InitSdkImage(""); err != nil {
		return nil, err
	}

	return &cli.Command{
		Name:      "image",
		Usage:     l.Get("image-usage"),
		ArgsUsage: "[image|-]",
		Aliases:   []string{"i"},
		Action:    imageAction,
		Flags:     imageFlags(),
	}, nil
}

func imageFlags() []cli.Flag {
	l := lang.GetLocalize()
	imageSdk := sdk.GetSdkImage()

	return []cli.Flag{
		&cli.StringFlag{
			Name:        "sdk",
			Aliases:     []string{"S"},
			Usage:       l.Get("sdk-usage"),
			DefaultText: imageSdk.GetName(),
			Category:    "global",
			Action: func(c *cli.Context, value string) error {
				if err := sdk.InitSdkImage(value); err != nil {
					return err
				}
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "model",
			Aliases:     []string{"m"},
			Usage:       l.Get("sdk-model-usage"),
			DefaultText: imageSdk.GetModel(),
			Category:    "image",
			Action: func(c *cli.Context, value string) error {
				imageSdk := sdk.GetSdkImage()
				imageSdk.SetModel(value)
				return nil
			},
		},
		&cli.StringFlag{
			Name:        "size",
			Aliases:     []string{"s"},
			Usage:       l.Get("image-size-usage"),
			DefaultText: imageSdk.GetSize(),
			Category:    "image",
			Action: func(c *cli.Context, value string) error {
				imageSdk := sdk.GetSdkImage()
				imageSdk.SetSize(value)
				return nil
			},
		},
		&cli.IntFlag{
			Name:        "image-nb",
			Aliases:     []string{"n"},
			Usage:       l.Get("image-nb-usage"),
			DefaultText: strconv.Itoa(imageSdk.GetImageNb()),
			Category:    "image",
			Action: func(c *cli.Context, value int) error {
				imageSdk := sdk.GetSdkImage()
				imageSdk.SetImageNb(value)
				return nil
			},
		},
		&cli.StringFlag{
			Name:     "output",
			Aliases:  []string{"o"},
			Usage:    l.Get("image-output-usage"),
			Category: "image",
			Action: func(c *cli.Context, value string) error {
				imageSdk := sdk.GetSdkImage()
				imageSdk.SetOutput(value)
				return nil
			},
		},
	}
}

func imageAction(c *cli.Context) error {
	return nil
}
