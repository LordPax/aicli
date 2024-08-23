package commands

import (
	"fmt"

	"github.com/LordPax/aicli/lang"
	"github.com/LordPax/aicli/service"
	"github.com/LordPax/aicli/utils"

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
	log, _ := utils.GetLog()

	if c.NArg() == 0 {
		return fmt.Errorf(l.Get("no-args"))
	}

	prompt := c.Args().First()
	response, err := service.SendTextRequest(prompt)
	if err != nil {
		return err
	}

	log.Printf(l.Get("hello-world"), response)

	return nil
}
