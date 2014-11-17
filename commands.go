package main

import (
	"github.com/mitchellh/cli"
	"github.com/wuub/roj/command"
	"os"
)

var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				Ui:                ui,
			}, nil
		},
		"agent": func() (cli.Command, error) {
			return &command.AgentCommand{
				Ui: ui,
			}, nil
		},
		"nodes": func() (cli.Command, error) {
			return &command.NodesCommand{
				Ui: ui,
			}, nil
		},
	}
}
