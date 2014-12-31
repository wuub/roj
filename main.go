package main

import (
	"flag"
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

	if args[0] == "apps" {
		appsFlags := flag.NewFlagSet("apps", flag.ContinueOnError)
		cluster := appsFlags.String("cluster", "http://127.0.0.1:8500", "consul cluster address")
		appsFlags.Parse(args[1:])

		fmt.Println(*cluster)

		cli, err := roj.NewClient(*cluster)
		if err != nil {
			log.Fatal(err)
		}
		apps, err := cli.Apps()
		if err != nil {
			log.Fatal(err)
		}
		for _, app := range apps {
			fmt.Printf("%s \n", app)
		}
		os.Exit(0)
	}

	os.Exit(1)
}
