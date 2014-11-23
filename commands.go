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
		"agent": func() (cli.Command, error) { return &command.AgentCommand{Ui: ui}, nil },
		"nodes": func() (cli.Command, error) { return &command.NodesCommand{Ui: ui}, nil },
	}
}
