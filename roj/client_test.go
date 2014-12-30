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
