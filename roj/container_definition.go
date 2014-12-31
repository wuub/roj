package roj

type ContainerDefinition struct {
	Image          string
	PublishedPorts []string
	Volumes        []string
}

func NewContainerDefinition() ContainerDefinition {
	return *new(ContainerDefinition)
}
