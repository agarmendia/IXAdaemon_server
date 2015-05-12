package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func komunikazioaKudeatu(conn net.Conn, in io.WriteCloser, out io.ReadCloser, errr io.ReadCloser) {
	for {

		//bezeroaren mezua jaso
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}

		//bezeroaren mezua exekutagarriari bidali
		_, err = in.Write([]byte(message))
		if err != nil {
			panic(err)
		}

		//exekutagarriarengandik mezua jaso
		a, err := bufio.NewReader(out).ReadString('\n')
		if err != nil {
			panic(err)
		}

		//mezu berria bezeroari bidali
		conn.Write([]byte(a + "\n"))
		if err != nil {
			return
		}
	}
}
