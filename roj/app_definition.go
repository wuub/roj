package roj

type AppDefinition struct {
	ID         string
	Name       string
	Version    string
	Containers map[string]ContainerDefinition
}
