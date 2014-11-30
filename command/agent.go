package command

import (
	"encoding/json"
	"net/http"

	"github.com/mitchellh/cli"
	"github.com/wuub/roj/roj"
)

type AgentCommand struct {
	Ui cli.Ui
}

func (c *AgentCommand) Help() string {
	return ""
}

type HTTPServer struct {
	Node *roj.Node
}

func (h *HTTPServer) HandleSysInfo(rw http.ResponseWriter, req *http.Request) {
	buf, _ := json.MarshalIndent(h.Node.SysInfo(), "", " ")
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(buf)
}

func (c *AgentCommand) Run(_ []string) int {
	node := roj.NewNode()
	err := roj.RegisterNode(*node)
	if err != nil {
		panic(err)
	}

	h := HTTPServer{Node: node}
	http.HandleFunc("/v1/node/sysinfo", h.HandleSysInfo)
	http.ListenAndServe("0.0.0.0:8000", nil)
	return 0
}
func (c *AgentCommand) Synopsis() string {
	return "Prints Roj nodes"
}
