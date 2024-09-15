package commands

import (
	"fmt"

	"github.com/LordPax/aicli/lang"
	"github.com/LordPax/aicli/service"

	cli "github.com/urfave/cli/v2"
)

func TextCommand() *cli.Command {
	l := lang.GetLocalize()
	return &cli.Command{
		Name:      "text",
		Usage:     l.Get("text-usage"),
		ArgsUsage: "[prompt]",
		Aliases:   []string{"t"},
		Action:    testAction,
	}
}

func testAction(c *cli.Context) error {
	l := lang.GetLocalize()

	if c.NArg() == 0 {
		return fmt.Errorf(l.Get("no-args"))
	}

	prompt := c.Args().First()
	if err := service.SendTextRequest(prompt); err != nil {
		return err
	}

	return nil
}
