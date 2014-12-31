package main

import (
	"fmt"
	"log"
	"os"

	"github.com/wuub/roj/roj"
)

const VERSION = "0.1.0"

func usage() {
	fmt.Printf("roj %s\n", VERSION)
	fmt.Println("    version    print roj version")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		usage()
		os.Exit(0)
	}

	if args[0] == "version" {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	urn := os.Getenv("ROJ_CONSUL")
	if urn == "" {
		urn = "http://127.0.0.1:8500"
	}
	cli, err := roj.NewClient(urn)
	if err != nil {
		log.Fatal(err)
	}

	switch args[0] {
	case "apps":
		apps, err := cli.Apps()
		if err != nil {
			log.Fatal(err)
		}
		for _, app := range apps {
			fmt.Printf("%s \n", app)
		}
		os.Exit(0)
	case "create":
		app := roj.NewAppDefinition()
		if err := app.Name.Set(args[1]); err != nil {
			log.Fatal(err)
		}

		container := roj.NewContainerDefinition()
		container.Image = args[2]

		app.Containers[app.ID+"_"+app.Name.Name] = container

		if err = cli.AddAppDefinition(app); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v", app)
		os.Exit(0)

	}

	os.Exit(1)
}
