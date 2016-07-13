package main

import (
	"fmt"
	"net"
	"os"
)

func initializeServer(mainPort string, ctrlPort string) (*net.TCPListener, *net.TCPListener) {

	dlog.Println("Launching server...\n")

	//Server configurations

	mainService := ":" + mainPort
	tcpAddr1, err := net.ResolveTCPAddr("tcp", mainService)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	ctrlService := ":" + ctrlPort
	tcpAddr2, err := net.ResolveTCPAddr("tcp", ctrlService)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	//listen on main Port
	ln, err := net.ListenTCP("tcp", tcpAddr1)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	//listen on control Port
	ln2, err := net.ListenTCP("tcp", tcpAddr2)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	return ln, ln2
}
