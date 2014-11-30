package template

import (
	"encoding/json"

	"github.com/armon/consul-api"
	"github.com/fsouza/go-dockerclient"
)

const templatePrefix = "roj/templates/"
const defaultSchemaVersion = "1.0.0"
const defaultTag = "latest"

type ContainerDefinition struct {
	Config     docker.Config     `json:"config,omitempty"`
	HostConfig docker.HostConfig `json:"host_config,omitempty"`
}

type Template struct {
	SchemaVersion string                         `json:"schema_version"`
	Name          string                         `json:"name,omitempty"`
	Tag           string                         `json:"tag,omitempty"`
	Containers    map[string]ContainerDefinition `json:"containers,omitempty"`
}

func (t *Template) SetDefaults() {
	if t.SchemaVersion == "" {
		t.SchemaVersion = defaultSchemaVersion
	}
	if t.Tag == "" {
		t.Tag = defaultTag
	}
}

func (t *Template) Upload(consul *consulapi.Client) error {
	t.SetDefaults()
	content, err := json.Marshal(t)
	if err != nil {
		return err
	}
	p := &consulapi.KVPair{Key: templatePrefix + t.Name + "/" + t.Tag, Value: content}
	_, err = consul.KV().Put(p, nil)
	return err
}

func (t *Template) String() string {
	content, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(content)
}
