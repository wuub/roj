package template

import (
	"encoding/json"

	"github.com/armon/consul-api"
	"github.com/fsouza/go-dockerclient"
)

const (
	templatePrefix       = "roj/templates/"
	defaultSchemaVersion = "1.0.0"
	defaultTag           = "latest"
)

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

func New(data []byte) (t *Template, err error) {
	if err = t.Unmarshal(data); err != nil {
		return
	}
	if err = t.Unmarshal(data); err != nil {
		return
	}
	return t, nil
}

func (t *Template) SetDefaults() {
	if t.SchemaVersion == "" {
		t.SchemaVersion = defaultSchemaVersion
	}
	if t.Tag == "" {
		t.Tag = defaultTag
	}
}

func (t *Template) Key() string {
	t.SetDefaults()
	return t.Name + "/" + t.Tag
}

func (t *Template) Upload(consul *consulapi.Client) error {
	content, err := json.Marshal(t)
	if err != nil {
		return err
	}
	p := &consulapi.KVPair{Key: templatePrefix + t.Key(), Value: content}
	_, err = consul.KV().Put(p, nil)
	return err
}

func (t *Template) Fetch(consul *consulapi.Client) error {
	pair, _, err := consul.KV().Get(templatePrefix+t.Key(), nil)
	if err != nil {
		return err
	}
	t.Unmarshal(pair.Value)
	return nil
}

func (t *Template) Unmarshal(data []byte) error {
	t.SetDefaults()
	return json.Unmarshal(data, t)
}

func (t *Template) String() string {
	content, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(content)
}

func ListTemplates(consul *consulapi.Client) (templates []Template, err error) {
	kvPairs, _, err := consul.KV().List(templatePrefix, nil)
	if err != nil {
		return
	}

	templates = make([]Template, len(kvPairs))

	for i, kvPair := range kvPairs {
		if err = templates[i].Unmarshal(kvPair.Value); err != nil {
			return
		}
	}

	return templates, nil
}
