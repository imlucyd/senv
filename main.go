package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/integrii/flaggy"
	"github.com/luisnquin/senv/cmd"
	"github.com/luisnquin/senv/log"
)

const DEFAULT_VERSION = "unversioned"

var (
	version = DEFAULT_VERSION
	commit  string

	//go:embed help.tpl
	helpTpl string
	//go:embed docs/senv.example.yaml
	genericConfigFile []byte
)

func main() {
	check := flaggy.NewSubcommand("check")
	check.Description = "Checks wether the current working directory has `senv.yaml` or `.env` files"
	flaggy.AttachSubcommand(check, 1)

	to := flaggy.NewSubcommand("to")
	to.Description = "Allows you to switch to other environment without a prompt, directly as an argument"
	flaggy.AttachSubcommand(to, 1)

	ls := flaggy.NewSubcommand("ls")
	ls.Description = "List all the environments in the found working directory."
	flaggy.AttachSubcommand(ls, 1)

	revert := flaggy.NewSubcommand("revert")
	revert.Description = "Reverts the last environment switch"
	flaggy.AttachSubcommand(revert, 1)

	init := flaggy.NewSubcommand("init")
	init.Description = "Creates a new configuration file in the current directory"
	flaggy.AttachSubcommand(init, 1)

	flaggy.SetName("senv")
	flaggy.SetDescription("Switch between .env files")
	flaggy.SetVersion(fmt.Sprintf("senv %s <%s>", version, commit))
	flaggy.DefaultParser.SetHelpTemplate(helpTpl)
	flaggy.Parse()

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	switch {
	case revert.Used:
		println("revert")
	case ls.Used:
		println("ls")
	case to.Used:
		println("to")
	case check.Used:
		if err := cmd.Check(); err != nil {
			log.Pretty.Fatal(err.Error())
		}
	case init.Used:
		if err := cmd.Init(genericConfigFile); err != nil {
			log.Pretty.Error(err.Error())
		}
	default:
		if err := cmd.Switch(currentDir); err != nil {
			log.Pretty.Error(err.Error())
		}
	}
}

		}
	}
}
