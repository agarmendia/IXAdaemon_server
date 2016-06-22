package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

func manageCommunication(conn net.Conn, ctrlPort string, command string, args []string) error {
	writech := make(chan bool)
	readch := make(chan bool)

	scanner := bufio.NewScanner(conn)
	bufout := bufio.NewWriter(conn)
	go writetonative(*scanner, lProcess.stdin, writech)
	go func() {
		checked := make(chan bool)
		go readfromnative(*bufout, lProcess.stdout, readch, checked)
		select {
		case <-checked:
			readch <- true
			return

		case <-time.After(10 * time.Second):
			readch <- false
			return
		}
	}()

	correctWriting := <-writech
	if correctWriting == false {
		return errors.New("Communication failure: Writing to Native")
	}

	fmt.Println("endwritetonative")
	correctReading := <-readch
	if correctReading == false {
		return errors.New("Communication failure: Writing to Native")
	}

	fmt.Println("endreadfromnative")
	return nil
}

func writetonative(scanner bufio.Scanner, in io.WriteCloser, c chan bool) {

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		message := scanner.Text()
		//fmt.Print("=======SEND=======")
		//fmt.Println(string(message))
		_, err := in.Write([]byte(message + "\n"))
		if err != nil {
			c <- false
		}
		if message == "[IXAdaemon]EOF" {
			//fmt.Println("manageCommunication has recived [IXAdaemon]EOF")
			c <- true
			break
		}
	}
	return
}

func readfromnative(bufout bufio.Writer, out io.ReadCloser, c chan bool, checked chan bool) {
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
	c <- true
	checked <- true
	return
}

func reconnect(conn net.Conn, ctrlPort string, command string, args []string) *LaunchedProcess {
	LaunchedProcess := launchNative(command, args)
	return LaunchedProcess
}
