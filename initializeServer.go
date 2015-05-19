package main

import (
	"fmt"
	"net"
)

func initializeServer(mainPort string, ctrlPort string) (*net.TCPListener, *net.TCPListener) {

	fmt.Println("Launching server...\n")

	//Server configurations

	mainService := ":" + mainPort
	tcpAddr1, err := net.ResolveTCPAddr("tcp", mainService)
	checkErrors(err)

	ctrlService := ":" + ctrlPort
	tcpAddr2, err := net.ResolveTCPAddr("tcp", ctrlService)
	checkErrors(err)

	//listen on main Port
	ln, err := net.ListenTCP("tcp", tcpAddr1)
	checkErrors(err)

	//listen on control Port
	ln2, err := net.ListenTCP("tcp", tcpAddr2)
	checkErrors(err)

	return ln, ln2
}
