package command

import (
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/wuub/roj/roj"
)

type NodesCommand struct {
	Ui cli.Ui
}

func (c *NodesCommand) Help() string {
	return ""
}
func (c *NodesCommand) Run(_ []string) int {

	nodes, _ := roj.FindNodes(false)

	for _, n := range nodes {
		fmt.Printf("%s\n", n.String())
	}
	return 0
}
func (c *NodesCommand) Synopsis() string {
	return "Prints Roj nodes"
}
