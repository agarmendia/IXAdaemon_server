package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
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
		if err != nil {
			*state = "3"
		} else {
			if message == "[IXAdaemon]INIT" {
				fmt.Println(message)
				*state = "1"
			}

			if message == "[IXAdaemon]RUN" {
				panic(5)
				*state = "0"
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
