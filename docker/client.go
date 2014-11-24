package docker

import "github.com/fsouza/go-dockerclient"

func NewDockerClient() (*docker.Client, error) {
	return docker.NewClient("unix:///var/run/docker.sock")
}

func FindInstances() (inatnces []string, err error) {
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
		if err == nil {
			continue
		}
		names[i] = inspectData.Image
	}

	return names, nil
}
