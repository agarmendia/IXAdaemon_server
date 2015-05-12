package main

import (
	"fmt"
	"net"
)

func zerbitzaria() *net.TCPListener {

	fmt.Println("Zerbitzaria abiarazten...\n")

	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	return ln
}
