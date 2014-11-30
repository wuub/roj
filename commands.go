package main

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/wuub/roj/command"
)

var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"agent":        func() (cli.Command, error) { return &command.AgentCommand{Ui: ui}, nil },
		"nodes":        func() (cli.Command, error) { return &command.NodesCommand{Ui: ui}, nil },
		"instances":    func() (cli.Command, error) { return &command.InstancesCommand{Ui: ui}, nil },
		"template":     func() (cli.Command, error) { return &command.TemplateCommand{Ui: ui}, nil },
		"launch":       func() (cli.Command, error) { return &command.LaunchCommand{Ui: ui}, nil },
		"local_launch": func() (cli.Command, error) { return &command.LocalLaunchCommand{Ui: ui}, nil },
	}
}
