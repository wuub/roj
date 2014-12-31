package roj

import (
	"errors"
	"strings"
)

type Client interface {
	Apps() (map[string]AppDefinition, error)
	AddAppDefinition(AppDefinition) error
}

func NewClient(urn string) (cli Client, err error) {

	if strings.HasPrefix(urn, "mock://") {
		return NewMockClient(urn)
	}
	return nil, errors.New("Unknown client urn: " + urn)
}
