package consul

import (
	"github.com/armon/consul-api"
	"github.com/wuub/roj/node"
)

func NewConsulClient() (*consulapi.Client, error) {
	return consulapi.NewClient(consulapi.DefaultConfig())
}

func RegisterNode(node node.Node) (err error) {
	client, err := NewConsulClient()
	if err != nil {
		return
	}

	agent := client.Agent()

	reg := &consulapi.AgentServiceRegistration{
		ID:   "roj-node",
		Name: "roj-node",
		Port: 8000,
	}

	if err := agent.ServiceRegister(reg); err != nil {
		return err
	}

	return nil
}
