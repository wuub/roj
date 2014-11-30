package roj

import (
	"encoding/json"
	"strings"

	"github.com/fsouza/go-dockerclient"
)

func NewDockerClient() (*docker.Client, error) {
	return docker.NewClient("unix:///var/run/docker.sock")
}

type Template struct {
	Id     string
	Config docker.Config
}

type Instance struct {
	Template
}

func FindInstances() (instances []string, err error) {
	client, err := NewDockerClient()
	if err != nil {
		return
	}

	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		return
	}

	names := make([]string, len(containers))
	for i, container := range containers {
		inspectData, err := client.InspectContainer(container.ID)
		if err != nil {
			continue
		}
		for _, envValue := range inspectData.Config.Env {
			if strings.Contains(envValue, "ROJ_ID") {
				by, _ := json.Marshal(inspectData)

				names[i] = string(by)
			}
		}
	}

	return names, nil
}

func LaunchTemplate(instanceId string, template *Template) error {
	return nil
}
