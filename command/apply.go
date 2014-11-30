package command

import (
	"log"
	"os"

	"github.com/armon/consul-api"
	"github.com/fsouza/go-dockerclient"
	"github.com/mitchellh/cli"
	"github.com/wuub/roj/node"
)

type ApplyCommand struct {
	Ui cli.Ui
}

func (c *ApplyCommand) Help() string {
	helpTest := `
`
	return helpTest
}
func (c *ApplyCommand) Run(_ []string) int {

	consulClient, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}

	nodeName, err := consulClient.Agent().NodeName()
	if err != nil {
		panic(err)
	}

	endpoint := os.Getenv("DOCKER_HOST")
	if endpoint == "" {
		endpoint = "unix:///var/run/docker.sock"
	}
	dockerClient, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}

	n, _ := node.New(consulClient, dockerClient, nodeName)

	err = n.ApplyAll()
	if err != nil {
		log.Fatal(err)
	}

	return 0
}
func (c *ApplyCommand) Synopsis() string {
	return "Manage Roj Templates"
}
