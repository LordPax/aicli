package commands

import (
	"clitemplate/lang"
	"clitemplate/utils"
	"fmt"

	cli "github.com/urfave/cli/v2"
)

func MainFlags() []cli.Flag {
	l := lang.GetLocalize()
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "silent",
			Aliases: []string{"s"},
			Usage:   l.Get("silent"),
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
	return fmt.Errorf(l.Get("no-command"))
}
