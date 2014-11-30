package roj

import "testing"

func TestNodeName(t *testing.T) {
	n := Node{Name: "hello"}

	if n.Name != "hello" {
		t.Error("node needs a name")
	}
}
