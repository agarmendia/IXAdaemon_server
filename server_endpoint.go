package main

import (
	"bufio"
	"fmt"
	"net"
)

type ServerEndpoint struct {
	conn   net.Conn
	output chan string
}

func newServerEndpoint(conn net.Conn) *ServerEndpoint {
	return &ServerEndpoint{
		conn:   conn,
		output: make(chan string)}
}

func (se *ServerEndpoint) Terminate() {
}

func (se *ServerEndpoint) Output() chan string {
	return se.output
}

func (se *ServerEndpoint) Send(msg string) bool {
	fmt.Println(msg)
	buf := bufio.NewWriter(se.conn)
	buf.WriteString(msg)
	buf.WriteString("\n")
	err := buf.Flush()
	if err != nil {
		return false
	}
	return true
}

func (se *ServerEndpoint) StartReading() {
	go se.read_client()
}

func (se *ServerEndpoint) read_client() {
	for {
		msg, err := bufio.NewReader(se.conn).ReadString('\n')
		if err != nil {
			break
		}
		se.output <- string(msg)
	}
	close(se.output)

}
