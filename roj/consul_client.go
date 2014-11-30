package roj

import (
	"encoding/json"

	"github.com/armon/consul-api"
)

func NewConsulClient() (*consulapi.Client, error) {
	return consulapi.NewClient(consulapi.DefaultConfig())
}

func RegisterNode(node Node) (err error) {
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

func FetchTemplate(name string) (template Template, err error) {
	client, err := NewConsulClient()
	if err != nil {
		return
	}
	kvPair, _, err := client.KV().Get(name, nil)
	if err != nil {
		return
	}

	template = Template{}
	err = json.Unmarshal(kvPair.Value, &template)
	if err != nil {
		return
	}

	return template, nil
}

func AssignTemplate(template, node string) (err error) {
	client, err := NewConsulClient()
	if err != nil {
		return
	}
	instanceId := "i-random"
	content, _ := json.Marshal(map[string]string{"id": instanceId, "template": template})
	p := &consulapi.KVPair{Key: "roj/instances/" + node + "/" + instanceId, Value: content}
	_, err = client.KV().Put(p, nil)
	return
}
