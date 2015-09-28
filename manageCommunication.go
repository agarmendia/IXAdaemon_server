package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	//"os"
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
		//fmt.Print("=======SEND=======")
		//fmt.Println(string(message))
		in.Write([]byte(message + "\n"))

		if message == "[IXAdaemon]EOF" {
			//fmt.Println("manageCommunication has recived [IXAdaemon]EOF")
			break
		}
	}
	c <- struct{}{}
	return
}

func readfromnative(bufout bufio.Writer, out io.ReadCloser, c chan struct{}) {
	sc := bufio.NewScanner(out)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		message := sc.Text()
		//fmt.Print("======RECEIVE=====")
		//fmt.Print(string(message))
		bufout.WriteString(message)
		bufout.WriteString("\n")
		bufout.Flush()
		if message == "[IXAdaemon]EOD" {
			//fmt.Println("manageCommunicationek has received [IXAdaemon]EOD")
			break
		}
	}
	c <- struct{}{}
	return
}
