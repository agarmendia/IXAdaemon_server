package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func manageCommunication(conn net.Conn, in io.WriteCloser, out io.ReadCloser, errr io.ReadCloser) {
	c := make(chan struct{})
	ch := make(chan struct{})
	scanner := bufio.NewScanner(conn)
	bufout := bufio.NewWriter(conn)
	go writetonative(*scanner, in, c)
	go readfromnative(*bufout, out, ch)
	<-c
	fmt.Println("endwritetonative")
	<-ch
	fmt.Println("endreadfromnative")

}

func writetonative(scanner bufio.Scanner, in io.WriteCloser, c chan struct{}) {

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		message := scanner.Text()
		in.Write([]byte(message + "\n"))

		if message == "[IXAdaemon]EOF" {
			fmt.Println("manageCommunicationek [IXAdaemon]EOF jaso du")
			break
		}
	}
	c <- struct{}{}
	return
}

func readfromnative(bufout bufio.Writer, out io.ReadCloser, c chan struct{}) {

	for {
		a, _, err := bufio.NewReader(out).ReadLine()
		if err != nil {
			fmt.Println(err)
		}

		_, err = bufout.WriteString(string(a))
		if err != nil {
			fmt.Println(err)
		}
		_, err = bufout.WriteString("\n")
		if err != nil {
			fmt.Println(err)
		}
		err = bufout.Flush()
		if err != nil {
			fmt.Println(err)
		}
		if string(a) == "[IXAdaemon]EOD" {
			fmt.Println("manageCommunicationek [IXAdaemon]EOD jaso du")
			break
		}
	}
	c <- struct{}{}
	return

}
