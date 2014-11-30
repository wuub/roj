package command

import (
	"flag"
	"fmt"

	"github.com/armon/consul-api"
	"github.com/mitchellh/cli"
	"github.com/wuub/roj/instance"
)

type createInstanceConf struct {
	template string
	node     string
}

type NewInstanceCommand struct {
	Ui   cli.Ui
	conf createInstanceConf
}

func (c *NewInstanceCommand) Help() string {
	helpTest := `
`
	return helpTest
}
func (c *NewInstanceCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("create-instance", flag.ContinueOnError)

	cmdFlags.StringVar(&c.conf.template, "template", "", "")
	cmdFlags.StringVar(&c.conf.node, "node", "", "")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	if c.conf.template == "" || c.conf.node == "" {
		c.Ui.Error("Borh -node and -template are mandatory options")
		return 1
	}

	consulClient, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}

	inst := instance.New(c.conf.node, c.conf.template)

	if err = inst.Upload(consulClient); err != nil {
		panic(err)
	}

	fmt.Printf("%s", inst.String())

	return 0
}
func (c *NewInstanceCommand) Synopsis() string {
	return "Manage Roj Templates"
}
