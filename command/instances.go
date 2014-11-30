package command

import (
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/wuub/roj/roj"
)

type InstancesCommand struct {
	Ui cli.Ui
}

func (c *InstancesCommand) Help() string {
	return ""
}
func (c *InstancesCommand) Run(_ []string) int {

	instances, _ := roj.FindInstances()

	for _, instance := range instances {
		fmt.Printf("%s\n", instance)
	}
	return 0
}
func (c *InstancesCommand) Synopsis() string {
	return "Prints local Roj instances"
}
