package main

import (
	"bufio"
	"io"
	"net"
)

func listenState(stdErr io.ReadCloser, listener *net.TCPListener) {
	var state string

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
		if err == nil {
			*state = message
		}
		if err == io.EOF {
			*state = "1"
			return
		}
	}

}

func writeState(conn net.Conn, state *string) {
	var b []byte
	for {
		conn.Read(b)
		_, err := conn.Write([]byte(*state))
		if err != nil {
			conn.Close()
			return
		}

	}
}
