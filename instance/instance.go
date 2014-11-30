package instance

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/armon/consul-api"
	"github.com/wuub/roj/template"
)

const (
	instancesPrefix = "roj/instances/"
)

type Instance struct {
	Id           string             `json:id`
	Node         string             `json:node`
	TemplateName string             `json:template`
	template     *template.Template `json:"-"`
	consulClient *consulapi.Client  `json:"-"`
}

func New(consul *consulapi.Client, node, templateName string) *Instance {
	inst := Instance{Node: node, TemplateName: templateName}

	id := make([]byte, 16)
	rand.Read(id)
	inst.Id = hex.EncodeToString(id)

	return &inst
}

func (i *Instance) Key() string {
	return i.Node + "/" + i.Id
}
func (i *Instance) Upload() error {
	content, err := json.Marshal(i)
	if err != nil {
		return err
	}
	p := &consulapi.KVPair{Key: instancesPrefix + i.Key(), Value: content}
	_, err = i.consulClient.KV().Put(p, nil)
	return err
}

func (i *Instance) String() string {
	res, _ := json.MarshalIndent(i, "", "  ")
	return string(res)
}

func (i *Instance) Template() *template.Template {
	if i.template == nil {
		parts := strings.Split(i.TemplateName, "/")
		i.template = &template.Template{Name: parts[0], Tag: parts[1]}
		if err := i.template.Fetch(i.consulClient); err != nil {
			panic(err)
		}
	}
	return i.template
}

func List(consul *consulapi.Client, prefix string) (instances []Instance, err error) {
	kvPairs, _, err := consul.KV().List(instancesPrefix+prefix, nil)
	if err != nil {
		return
	}

	instances = make([]Instance, len(kvPairs))
	for i, kvPair := range kvPairs {
		if err = json.Unmarshal(kvPair.Value, &instances[i]); err != nil {
			return
		}
		instances[i].consulClient = consul
	}

	return instances, nil
}
