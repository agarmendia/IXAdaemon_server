package main

import (
	"flag"
	"fmt"
	"os"
)

func parseArguments() (string, string, string, []string, string) {

	var mainPort string
	flag.StringVar(&mainPort, "mainPort", "2101", "aplication port")

	var ctrlPort string
	flag.StringVar(&ctrlPort, "ctrlPort", "2102", "control port")

	var logFile string
	flag.StringVar(&logFile, "log", "", "Set log file")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: IXAdaemon_server [--mainPort= 2101] [--ctrlPort = 2102] [--log = logfile.txt] \"command\"")
		fmt.Println("--mainPort: Port number for client communication")
		fmt.Println("--ctrlPort: Port number for control communication")
		fmt.Println("--log: If you pass a valid filename, will print logs in that page. If not, will print them on stdin")
		os.Exit(1)
	}

	command := args[0]
	args = args[1:]

	return mainPort, ctrlPort, command, args, logFile

}
