package roj

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"strings"
)

type AppName struct {
	Name    string
	Version string
}

func (a *AppName) Set(s string) error {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return errors.New("bad AppName")
	}
	a.Name = parts[0]
	a.Version = parts[1]
	return nil
}

type AppDefinition struct {
	ID         string
	Name       AppName
	Containers map[string]ContainerDefinition
}

func NewAppDefinition() AppDefinition {
	app := AppDefinition{}
	app.ID = randomString(16)
	app.Containers = make(map[string]ContainerDefinition)
	return app
}

func randomString(bytes int) string {
	b := make([]byte, bytes)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
