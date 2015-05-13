package main

import (
	"flag"
)

func argumentuakParseatu() (string, []string) {

	var portua string
	flag.StringVar(&portua, "portua", "2101", "aplikazioaren portua")

	flag.Parse()

	args := flag.Args()

	return portua, args

}
