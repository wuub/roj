package command

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/mitchellh/cli"
	"github.com/wuub/roj/roj"
)

type TemplateCommand struct {
	Ui cli.Ui
}

func (c *TemplateCommand) Help() string {
	return ""
}
func (c *TemplateCommand) Run(args []string) int {

	cmdFlags := flag.NewFlagSet("template", flag.ContinueOnError)

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	args = cmdFlags.Args()
	fmt.Print(args)
	if len(args) != 2 {
		c.Ui.Error("you need to give an action and a file_name")
		return 1
	}

	content, err := ioutil.ReadFile(args[1])
	if err != nil {
		return 2
	}

	err = roj.UploadTemplate(args[1], content)
	if err != nil {
		panic(err)
	}
	return 0
}
func (c *TemplateCommand) Synopsis() string {
	return "Manage Roj Templates"
}
