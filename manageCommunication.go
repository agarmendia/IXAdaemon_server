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
	commch := make(chan bool)

	scanner := bufio.NewScanner(conn)
	bufout := bufio.NewWriter(conn)
	go writetonative(*scanner, lProcess.stdin, writech, commch)

	//checked := make(chan bool)
	go readfromnative(*bufout, lProcess.stdout, readch, commch, writech)

	correctWriting := <-writech
	if correctWriting == false {
		return errors.New("Communication failure: Writing to Native")
		panic(10)
	}

	fmt.Println("endwritetonative")
	correctReading := <-readch
	if correctReading == false {
		return errors.New("Communication failure: Writing to Native")
	}

	fmt.Println("endreadfromnative")
	return nil
}

func writetonative(scanner bufio.Scanner, in io.WriteCloser, c chan bool, commch chan bool) {

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		message := scanner.Text()
		//fmt.Print("=======SEND=======")
		//fmt.Println(string(message))
		_, err := in.Write([]byte(message + "\n"))
		if err != nil {
			c <- false
			break
		}
		if message == "[IXAdaemon]EOF" {
			//fmt.Println("manageCommunication has recived [IXAdaemon]EOF")
			c <- true
			commch <- true
			break
		}
	}
	return
}

func readfromnative(bufout bufio.Writer, out io.ReadCloser, c chan bool, commch chan bool, writech chan bool) {
	<-commch
	sc := bufio.NewScanner(out)
	var message string
	sc.Split(bufio.ScanLines)
	for true {
		//fmt.Println("eii")
		ch := make(chan bool)
		go func() {
			sc.Scan()
			//if err != nil {
			//	ch <- false
			//} else {
			ch <- true
			//}
			return
		}()
		select {
		case correct := <-ch:
			if correct == true {
				message = sc.Text()
			} else {
				c <- false
				return
			}
		case <-time.After(time.Second * 10):
			fmt.Println("TIMEOUT")
			writech <- true
			c <- false
			return
		}

		//message := sc.Text()
		//fmt.Print("======RECEIVE=====")
		//fmt.Println(string(message))
		bufout.WriteString(message)
		bufout.WriteString("\n")
		bufout.Flush()
		if message == "[IXAdaemon]EOD" {
			//fmt.Println("manageCommunicationek has received [IXAdaemon]EOD")
			break
		}
	}
	c <- true
	//checked <- true
	return
}

func reconnect(conn net.Conn, ctrlPort string, command string, args []string) *LaunchedProcess {
	LaunchedProcess := launchNative(command, args)
	return LaunchedProcess
}
