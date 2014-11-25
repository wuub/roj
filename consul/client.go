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

func FindNodes(all bool) (nodes []*RemoteNode, err error) {
	client, err := NewConsulClient()

	if err != nil {
		return
	}

	services, _, err := client.Health().Service("roj-node", "", false, nil)

	if err != nil {
		return
	}

	nodes = make([]*RemoteNode, len(services))
	for idx, s := range services {
		nodes[idx] = &RemoteNode{s, nil}
	}

	return nodes, nil
}

func UploadTemplate(name string, content []byte) (err error) {
	client, err := NewConsulClient()
	if err != nil {
		return
	}
	p := &consulapi.KVPair{Key: "roj/templates/" + name, Value: content}
	_, err = client.KV().Put(p, nil)
	return
}
