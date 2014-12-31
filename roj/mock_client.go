package roj

type MockClient struct {
	apps map[string]AppDefinition
}

func (m *MockClient) Apps() (map[string]AppDefinition, error) {
	return m.apps, nil
}

func NewMockClient(urn string) (Client, error) {
	cli := new(MockClient)
	cli.apps = make(map[string]AppDefinition)
	return cli, nil
}

func (m *MockClient) AddAppDefinition(app AppDefinition) error {
	m.apps[app.ID] = app
	return nil
}
