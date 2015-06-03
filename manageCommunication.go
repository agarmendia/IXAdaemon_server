package main

import (
	"bufio"
	"io"
	"net"
)

func manageCommunication(conn net.Conn, in io.WriteCloser, out io.ReadCloser, errr io.ReadCloser) {

	//to modify

	for {

		//Get message from client
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}

		//Send message to native program
		_, err = in.Write([]byte(message))
		if err != nil {
			panic(err)
		}

		//Get message from native program
		a, err := bufio.NewReader(out).ReadString('\n')
		if err != nil {
			panic(err)
		}

		//Send message to client
		conn.Write([]byte(a + "\n"))
		if err != nil {
			return
		}

	}
}
