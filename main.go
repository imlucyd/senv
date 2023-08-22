package main

import (
	"os"
	"runtime/debug"

	"github.com/integrii/flaggy"
	"github.com/luisnquin/senv/internal"
	"github.com/luisnquin/senv/internal/assets"
	"github.com/luisnquin/senv/internal/cmd"
	"github.com/luisnquin/senv/internal/log"
)

var (
	version = internal.DEFAULT_VERSION
	commit  string
)

func main() {
	var toSwitchArg string

	to := flaggy.NewSubcommand("to")
	to.Description = "Allows you to switch to other environment without a prompt"
	flaggy.AttachSubcommand(to, 1)
	to.AddPositionalValue(&toSwitchArg, "environment", 1, true, "I don't know what this does")

	ls := flaggy.NewSubcommand("ls")
	ls.Description = "List all the environments in the working directory"
	flaggy.AttachSubcommand(ls, 1)

	init := flaggy.NewSubcommand("init")
	init.Description = "Creates a new configuration file in the current directory"
	flaggy.AttachSubcommand(init, 1)

	out := flaggy.NewSubcommand("out")
	out.Description = "Print your current environment to stdout, exits with status code 1 if it can't"
	flaggy.AttachSubcommand(out, 1)

	var completionShellArg string

	completion := flaggy.NewSubcommand("completion")
	completion.Hidden = true
	flaggy.AttachSubcommand(completion, 1)
	completion.AddPositionalValue(&completionShellArg, "shell", 1, true, "Supported shells: zsh && bash")

	var versionFlag bool

	flaggy.DefaultParser.ShowVersionWithVersionFlag = false // Thanks but ugly
	flaggy.Bool(&versionFlag, "v", "version", "Displays the program version.")

	flaggy.SetName(internal.PROGRAM_NAME)
	flaggy.SetDescription("Switch your .env file")
	flaggy.DefaultParser.SetHelpTemplate(assets.GetHelpTpl())
	flaggy.Parse()

	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	switch {
	case versionFlag:
		if err := cmd.Version(version, commit); err != nil {
			log.Pretty.Fatal(err.Error())
		}
	case completion.Used:
		if err := cmd.Completion(completionShellArg); err != nil {
			log.Pretty.Error(err.Error())
		}
	case out.Used:
		if err := cmd.Out(); err != nil {
			log.Pretty.Fatal(err.Error())
		}
	case ls.Used:
		if err := cmd.Ls(currentDir); err != nil {
			log.Pretty.Error(err.Error())
		}
	case to.Used:
		if err := cmd.SwitchTo(currentDir, toSwitchArg); err != nil {
			log.Pretty.Error(err.Error())
		}
	case init.Used:
		if err := cmd.Init(); err != nil {
			log.Pretty.Error(err.Error())
		}
	default:
		if err := cmd.Switch(currentDir); err != nil {
			log.Pretty.Error(err.Error())
		}
	}
}

func getCommit() string {
	info, ok := debug.ReadBuildInfo()
	if ok {
		for _, kv := range info.Settings {
			if kv.Key == "vcs.revision" {
				return kv.Value
			}
		}
	}

	return commit
}
