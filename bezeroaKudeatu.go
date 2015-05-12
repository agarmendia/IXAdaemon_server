package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func bezeroaKudeatu(conn net.Conn, in io.WriteCloser, out io.ReadCloser, errr io.ReadCloser) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		fmt.Print("Bidaltzeko mezua:", string(message))
		// sample process for string received
		_, err = in.Write([]byte(message))
		if err != nil {
			panic(err)
		}
		a, err := bufio.NewReader(out).ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Println("aplikaziotik jasotako mezua: ", a)
		// send new string back to client
		conn.Write([]byte(a + "\n"))
		if err != nil {
			return
		}
	}
}
