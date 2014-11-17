package command

import (
	"github.com/armon/consul-api"
	"github.com/mitchellh/cli"
)

type NodesCommand struct {
	Ui cli.Ui
}

func (c *NodesCommand) Help() string {
	return ""
}
func (c *NodesCommand) Run(_ []string) int {
	config := consulapi.DefaultConfig()
	config.Address = "172.16.1.1:80"
	client, _ := consulapi.NewClient(config)
	catalog := client.Catalog()

	nodes, _, err := catalog.Nodes(nil)
	if err != nil {
		panic(err)
	}
	for _, node := range nodes {
		c.Ui.Output(node.Node + node.Address)
	}
	return 0
}
func (c *NodesCommand) Synopsis() string {
	return "Prints Roj nodes"
}
