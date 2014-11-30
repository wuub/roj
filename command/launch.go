package command

import (
	"flag"
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/wuub/roj/roj"
)

type LaunchCommand struct {
	Ui cli.Ui
}

func (c *LaunchCommand) Help() string {
	return "launch [template] [node]"
}
func (c *LaunchCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("launch", flag.ContinueOnError)

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	args = cmdFlags.Args()
	fmt.Print(args)
	if len(args) != 2 {
		c.Ui.Error("you need to give a template and node")
		return 1
	}
	template := args[0]
	node := args[1]

	err := roj.LaunchTemplate(template, node)
	if err != nil {
		panic(err)
	}
	return 0
}
func (c *LaunchCommand) Synopsis() string {
	return "Launch Roj Templates"
}
