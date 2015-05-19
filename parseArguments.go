package main

import (
	"flag"
	"fmt"
	"os"
)

func parseArguments() (string, string, []string) {

	var mainPort string
	flag.StringVar(&mainPort, "mainPort", "2101", "aplication port")

	var ctrlPort string
	flag.StringVar(&ctrlPort, "ctrlPort", "2102", "control port")

	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println("Usage: wrapper (--mainPort= 2101) (--ctrlPort = 2102) \"language\" \"native program\"")
		os.Exit(1)
	}

	return mainPort, ctrlPort, args

}
