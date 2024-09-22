package main

import (
	"fmt"
	"os"

	"github.com/LordPax/aicli/commands"
	"github.com/LordPax/aicli/config"
	"github.com/LordPax/aicli/lang"
	"github.com/LordPax/aicli/sdk"
	"github.com/LordPax/aicli/utils"

	cli "github.com/urfave/cli/v2"
	ini "gopkg.in/ini.v1"
)

func main() {
	if err := config.InitConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	log, err := utils.GetLog()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer log.Close()

	config.CONFIG_INI, err = ini.Load(config.CONFIG_FILE)
	if err != nil {
		log.PrintfErr("%v\n", err)
		os.Exit(1)
	}

	l := lang.GetLocalize()
	l.SetLang(os.Getenv("LANG"))
	l.AddStrings(&lang.EN_STRINGS, "en_US.UTF-8", "en_GB.UTF-8", "en")
	l.AddStrings(&lang.FR_STRINGS, "fr_FR.UTF-8", "fr_CA.UTF-8", "fr")

	if err := sdk.InitSdk(); err != nil {
		log.PrintfErr("%v\n", err)
		os.Exit(1)
	}

	app := cli.NewApp()
	app.Name = config.NAME
	app.Usage = l.Get("usage")
	app.Version = config.VERSION
	app.Action = commands.MainAction
	app.Flags = commands.MainFlags()
	app.Commands = []*cli.Command{
		commands.TextCommand(),
		commands.TranslateCommand(),
		// TODO : add command for image, audio and translate
	}

	if err := app.Run(os.Args); err != nil {
		log.PrintfErr("%v\n", err)
	}

	_ = utils.RmTmpDir()
}
