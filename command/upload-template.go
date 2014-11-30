package command

import (
	"flag"
	"io/ioutil"

	"github.com/armon/consul-api"
	"github.com/mitchellh/cli"
	"github.com/wuub/roj/template"
)

type uploadConf struct {
	file string
}

type UploadTemplateCommand struct {
	Ui   cli.Ui
	conf uploadConf
}

func (c *UploadTemplateCommand) Help() string {
	helpTest := `
Usage: roj upload-template [options]

Options:
   -file="-"       File to read template from or put template to. default: "-" 
`
	return helpTest
}
func (c *UploadTemplateCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("template", flag.ContinueOnError)
	cmdFlags.StringVar(&c.conf.file, "file", "-", "")

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	content, err := ioutil.ReadFile(c.conf.file)
	if err != nil {
		c.Ui.Error("Must specify an existing template file")
		c.Ui.Error("")
		return 1
	}

	consulClient, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		panic(err)
	}

	t, err := template.New(consulClient, content)
	if err != nil {
		c.Ui.Error("Must specify a valid template file")
		c.Ui.Error("")
		return 1
	}

	err = t.Upload()
	if err != nil {
		panic(err)
	}

	c.Ui.Output(t.Name)
	return 0
}
func (c *UploadTemplateCommand) Synopsis() string {
	return "Manage Roj Templates"
}
