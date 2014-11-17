package command

import (
	"github.com/armon/consul-api"
	"github.com/mitchellh/cli"
	"time"
)

type AgentCommand struct {
	Ui cli.Ui
}

func (c *AgentCommand) Help() string {
	return ""
}
func (c *AgentCommand) Run(_ []string) int {
	config := consulapi.DefaultConfig()
	config.Address = "172.16.1.1"
	client, _ := consulapi.NewClient(config)
	agent := client.Agent()

	reg := &consulapi.AgentServiceRegistration{
		ID:   "roj-node",
		Name: "roj-node",
		Port: 8000,
	}

	if err := agent.ServiceRegister(reg); err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)
	defer agent.ServiceDeregister(reg.ID)

	return 0
}
func (c *AgentCommand) Synopsis() string {
	return "Prints Roj nodes"
}
