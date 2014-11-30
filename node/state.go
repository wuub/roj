package node

import (
	"log"
	"strings"

	"github.com/armon/consul-api"
	"github.com/fsouza/go-dockerclient"
	"github.com/wuub/roj/instance"
	"github.com/wuub/roj/template"
)

type InstanceState struct {
	Instance           instance.Instance
	Template           *template.Template
	ExistingConatiners map[string]ContainerState
	RequiredContainers map[string]ContainerState
	dockerClient       *docker.Client
	consulClient       *consulapi.Client
}

type ContainerState struct {
	ID string
}

func envMap(env []string) map[string]string {
	out := make(map[string]string, len(env))
	for _, e := range env {
		parts := strings.SplitN(e, "=", 2)
		out[parts[0]] = parts[1]
	}
	return out
}

func NewInstanceState(dockerClient *docker.Client, consulClient *consulapi.Client, instance instance.Instance) (is InstanceState, err error) {
	is = InstanceState{Instance: instance, Template: instance.Template(), dockerClient: dockerClient, consulClient: consulClient}
	is.ExistingConatiners = map[string]ContainerState{}
	is.RequiredContainers = map[string]ContainerState{}

	for name, def := range is.Template.Containers {
		img, err := dockerClient.InspectImage(def.Config.Image)
		if err != nil {
			return is, err
		}
		is.RequiredContainers[name] = ContainerState{ID: img.ID}
	}

	for name, container := range is.Containers() {
		is.ExistingConatiners[name] = ContainerState{ID: container.Image}
	}

	return is, nil
}

func (i *InstanceState) Containers() map[string]docker.Container {
	opts := docker.ListContainersOptions{All: true}
	apiContainers, err := i.dockerClient.ListContainers(opts)
	if err != nil {
		panic(err)
	}
	res := make(map[string]docker.Container, 0)

	for _, apiContainer := range apiContainers {
		container, err := i.dockerClient.InspectContainer(apiContainer.ID)
		if err != nil {
			panic(err)
		}
		env := envMap(container.Config.Env)
		if env["ROJ_ID"] == i.Instance.Id {
			res[env["ROJ_NAME"]] = *container
		}
	}
	return res
}

func (i *InstanceState) RebuildRequired() bool {
	if len(i.ExistingConatiners) != len(i.RequiredContainers) {
		log.Printf("Rebuild required [%s %s]. Different number of containers", i.Template.Key(), i.Instance.Id)
		return true
	}

	for k, required := range i.RequiredContainers {
		existing := i.ExistingConatiners[k]
		if required != existing {

			log.Printf("Rebuild required [%s %s]: Container state changed: [%s] %v != %v",
				i.Template.Key(), i.Instance.Id, k, required, existing)
			return true
		}
	}
	return false
}
