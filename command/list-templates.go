package command

import (
	"flag"
	"fmt"

	"github.com/armon/consul-api"
	"github.com/mitchellh/cli"
	"github.com/wuub/roj/template"
)

type ListTemplatesCommand struct {
	Ui cli.Ui
}

func (c *ListTemplatesCommand) Help() string {
	helpTest := `
Usage: roj list-templates [options]

`
	return helpTest
}
func (c *ListTemplatesCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("list-templates", flag.ContinueOnError)

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	consulClient, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}

	templates, err := template.ListTemplates(consulClient)
	if err != nil {
		panic(err)
	}

	for _, t := range templates {
		c.Ui.Output(fmt.Sprintf("%s/%s\t%d", t.Name, t.Tag, len(t.Containers)))
	}

	return 0
}
func (c *ListTemplatesCommand) Synopsis() string {
	return "Manage Roj Templates"
}
