package node

import (
	"fmt"
	"log"

	"github.com/armon/consul-api"
	"github.com/fsouza/go-dockerclient"
	"github.com/wuub/roj/instance"
)

type Node struct {
	Name         string
	Instances    map[string]*instance.Instance
	Containers   map[string]*docker.Container
	consulClient *consulapi.Client
	dockerClient *docker.Client
}

func New(consulClient *consulapi.Client, dockerClient *docker.Client, nodeName string) (*Node, error) {
	n := new(Node)
	n.Name = nodeName
	n.consulClient = consulClient
	n.dockerClient = dockerClient
	return n, nil
}

func (n *Node) ApplyAll() (err error) {
	instances, err := instance.List(n.consulClient, n.Name)
	if err != nil {
		return
	}

	if err = n.PullImages(&instances); err != nil {
		return
	}

	toUpdate := []InstanceState{}
	for _, instance := range instances {
		is, _ := NewInstanceState(n.dockerClient, n.consulClient, instance)
		if is.RebuildRequired() {
			toUpdate = append(toUpdate, is)
		}
	}

	for _, is := range toUpdate {
		log.Print(is.Template.Name)

		for _, cont := range is.Containers() {
			log.Printf("Removing: %s", cont.ID)
			opts := docker.RemoveContainerOptions{ID: cont.ID, Force: true}
			err := n.dockerClient.RemoveContainer(opts)
			if err != nil {
				panic(err)
			}
		}

		for name, def := range is.Template.Containers {
			def.Config.Env = append(def.Config.Env, "ROJ_ID="+is.Instance.Id)
			def.Config.Env = append(def.Config.Env, "ROJ_NAME="+name)
			opts := docker.CreateContainerOptions{Name: is.Instance.Id + "_" + name, Config: &def.Config}
			cont, err := n.dockerClient.CreateContainer(opts)
			if err != nil {
				panic(err)
			}
			fmt.Println(cont)
		}
	}

	return nil
}

func (n *Node) PullImages(instances *[]instance.Instance) error {
	for _, instance := range *instances {
		template := instance.Template()
		for _, c := range template.Containers {
			opts := docker.PullImageOptions{Repository: c.Config.Image}
			log.Printf("Pulling image: %s for template: %s for instance: %s", opts.Repository, template.Key(), instance.Id)
			err := n.dockerClient.PullImage(opts, docker.AuthConfiguration{})
			if err != nil {
				return err
			}
		}
	}
	return nil
}
