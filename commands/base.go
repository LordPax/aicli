package commands

import (
	"errors"

	"github.com/LordPax/aicli/lang"
	"github.com/LordPax/aicli/utils"

	cli "github.com/urfave/cli/v2"
)

func MainFlags() []cli.Flag {
	l := lang.GetLocalize()
	return []cli.Flag{
		&cli.BoolFlag{
			Name:               "silent",
			Aliases:            []string{"s"},
			Usage:              l.Get("silent"),
			DisableDefaultText: true,
			Action: func(c *cli.Context, value bool) error {
				log, err := utils.GetLog()
				if err != nil {
					return err
				}

				log.SetSilent(value)

				return nil
			},
		},
	}
}

func MainAction(c *cli.Context) error {
	l := lang.GetLocalize()
	return errors.New(l.Get("no-command"))
}
