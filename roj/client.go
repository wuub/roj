package roj

import (
	"errors"
	"strings"
)

type App struct {
}

type Client interface {
	Apps() ([]App, error)
}

func NewClient(urn string) (cli Client, err error) {

	if strings.HasPrefix(urn, "mock://") {
		return NewMockClient(urn)
	}
	return nil, errors.New("Unknown client urn: " + urn)
}
