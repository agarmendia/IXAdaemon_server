package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func listenState(stdErr io.ReadCloser, listener *net.TCPListener) {
	var state string
	state = "3"
	for {

		conn, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go askState(stdErr, &state)
		go writeState(conn, &state)

	}

}

func askState(stdErr io.ReadCloser, state *string) {

	for {
		message, err := bufio.NewReader(stdErr).ReadString('\n')
		fmt.Println("Hau da estadoa: " + *state)
		if err != nil {
			fmt.Println(err)
			*state = "3"
		} else {
			if strings.Contains(message, "[IXAdaemon]INIT") {
				//panic(1)
				fmt.Println(message)
				*state = "1"
			} else {
				if strings.Contains(message, "[IXAdaemon]RUN") {
					//panic(2)
					*state = "0"
					//break
				}
			}

		}
	}
}

func writeState(conn net.Conn, state *string) {

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}

		if message == "state\n" {

			conn.Write([]byte(*state + "\n"))
		}

	}

}
