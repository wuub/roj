package instance

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"

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
}

func New(node, templateName string) *Instance {
	inst := Instance{Node: node, TemplateName: templateName}

	id := make([]byte, 16)
	rand.Read(id)
	inst.Id = hex.EncodeToString(id)

	return &inst
}

func (i *Instance) Key() string {
	return i.Node + "/" + i.Id
}
func (i *Instance) Upload(consul *consulapi.Client) error {
	content, err := json.Marshal(i)
	if err != nil {
		return err
	}
	p := &consulapi.KVPair{Key: instancesPrefix + i.Key(), Value: content}
	_, err = consul.KV().Put(p, nil)
	return err
}

func (i *Instance) String() string {
	res, _ := json.MarshalIndent(i, "", "  ")
	return string(res)
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
	}

	return instances, nil
}
