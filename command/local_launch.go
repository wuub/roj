package command

import (
	"flag"
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/wuub/roj/roj"
)

type LocalLaunchCommand struct {
	Ui cli.Ui
}

func (c *LocalLaunchCommand) Help() string {
	return "local_launch [instance-id] [template]"
}
func (c *LocalLaunchCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("local_launch", flag.ContinueOnError)

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	args = cmdFlags.Args()
	fmt.Print(args)
	if len(args) != 2 {
		c.Ui.Error("you need to give a template and node")
		return 1
	}
	instanceId := args[0]
	templateName := args[1]

	template, err := roj.FetchTemplate(templateName)
	if err != nil {
		panic(err)
	}

	err = roj.LaunchTemplate(instanceId, &template)
	if err != nil {
		panic(err)
	}
	return 0
}
func (c *LocalLaunchCommand) Synopsis() string {
	return "Launch Roj Templates"
}
