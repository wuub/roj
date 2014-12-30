package roj

type MockClient struct {
	apps []App
}

func (m *MockClient) Apps() ([]App, error) {
	return m.apps, nil
}

func NewMockClient(urn string) (Client, error) {
	cli := new(MockClient)
	return cli, nil
}
