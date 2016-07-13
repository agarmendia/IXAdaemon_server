package main

import (
	"bufio"
	"errors"
	"io"
	"net"
	"time"
)

func manageCommunication(conn net.Conn, ctrlPort string, command string, args []string) (error, []string) {
	writech := make(chan bool)
	readch := make(chan bool)
	commch := make(chan bool)
	currentDoc := []string{}

	scanner := bufio.NewScanner(conn)
	bufout := bufio.NewWriter(conn)
	go writetonative(*scanner, lProcess.stdin, writech, commch, &currentDoc)

	go readfromnative(*bufout, lProcess.stdout, readch, commch, writech)

	correctWriting := <-writech
	if correctWriting == false {
		return errors.New("Communication failure: Writing to Native"), currentDoc
		panic(10)
	}

	correctReading := <-readch

	if correctReading == false {
		return errors.New("Communication failure: Writing to Native"), currentDoc
	}

	return nil, nil
}

func writetonative(scanner bufio.Scanner, in io.WriteCloser, c chan bool, commch chan bool, currentDoc *[]string) {

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		message := scanner.Text()
		*currentDoc = append(*currentDoc, message) //may be useful for not losing information.

		_, err := in.Write([]byte(message + "\n"))
		if err != nil {
			c <- false
			break
		}
		if message == "[IXAdaemon]EOF" {

			c <- true
			commch <- true
			break
		}
	}
	return
}

func readfromnative(bufout bufio.Writer, out io.ReadCloser, c chan bool, commch chan bool, writech chan bool) {

	//waits until writetonative finishes (to manage timeouts)
	<-commch
	sc := bufio.NewScanner(out)
	var message string
	sc.Split(bufio.ScanLines)
	for true {

		ch := make(chan bool)
		go func() {
			sc.Scan()
			ch <- true
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
		case <-time.After(time.Minute * 10):
			dlog.Println("TIMEOUT")
			c <- false
			return
		}

		bufout.WriteString(message)
		bufout.WriteString("\n")
		bufout.Flush()
		if message == "[IXAdaemon]EOD" {

			break
		}
	}
	c <- true
	return
}
