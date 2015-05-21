package main

import (
	"bufio"
	"io"
	"net"
)

func listenState(stdErr io.ReadCloser, listener *net.TCPListener) {

	conn, err := listener.Accept()

	if err != nil {
		panic(err)
	}

	for {
		state, err := bufio.NewReader(stdErr).ReadString('\n')
		if err != nil {
			panic(err)
		}

		conn.Write([]byte(state + "\n"))

	}

}
