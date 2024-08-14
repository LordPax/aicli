package commands

import (
	"clitemplate/lang"
	"clitemplate/utils"
	"fmt"

	cli "github.com/urfave/cli/v2"
)

func TestCommand() *cli.Command {
	l := lang.GetLocalize()
	return &cli.Command{
		Name:      "test",
		Usage:     l.Get("test-usage"),
		ArgsUsage: "<name>",
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

	name := c.Args().First()
	log.Printf(l.Get("hello-world"), name)

	return nil
}
