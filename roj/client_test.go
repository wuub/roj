package roj

import "testing"

func TestClient_NewClient(t *testing.T) {
	cli, err := NewClient("mock://")
	if err != nil {
		t.Error(err)
	}

	if cli == nil {
		t.Error("no mock client returned")
	}
}

func TestClient_AddApp(t *testing.T) {
	cli, _ := NewClient("mock://")

	app := AppDefinition{ID: "test-id"}

	cli.AddAppDefinition(app)

	if apps, _ := cli.Apps(); len(apps) != 1 {
		t.Error("App Definition not added")
	}
}
