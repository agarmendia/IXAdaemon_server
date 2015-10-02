package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	//"os"
	"time"
)

func manageCommunication(conn net.Conn, in io.WriteCloser, out io.ReadCloser, errr io.ReadCloser, ctrlPort string) {
	writech := make(chan struct{})
	readch := make(chan struct{})
	//timeoutch := make(chan error, 1)

	scanner := bufio.NewScanner(conn)
	bufout := bufio.NewWriter(conn)
	go writetonative(*scanner, in, writech)
	go func() {
		checked := make(chan bool)
		go readfromnative(*bufout, out, readch, checked)
		select {
		case <-checked:
			return
		case <-time.After(time.Minute):
			bufout.WriteString("vamoh a plobal")
			bufout.WriteString("\n")
			bufout.Flush()
			//net.Dial("tcp", "127.0.0.1:"+ctrlPort)

		}
	}()
	//go readfromnative(*bufout, out, readch)
	<-writech
	fmt.Println("endwritetonative")
	<-readch
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

func readfromnative(bufout bufio.Writer, out io.ReadCloser, c chan struct{}, checked chan bool) {
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
	checked <- true
	return
}
