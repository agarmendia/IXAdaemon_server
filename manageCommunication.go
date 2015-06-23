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
		fmt.Print("======BIDALI======")
		fmt.Println(string(message))
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
	sc := bufio.NewScanner(out)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		message := sc.Text()
		fmt.Print("=======JASO=======")
		fmt.Println(string(message))
		bufout.WriteString(message)
		bufout.WriteString("\n")
		bufout.Flush()
		if message == "[IXAdaemon]EOD" {
			fmt.Println("manageCommunicationek [IXAdaemon]EOD jaso du")
			break
		}
	}
	c <- struct{}{}
	return
}

/*func readfromnative(bufout bufio.Writer, out io.ReadCloser, c chan struct{}) {
	//var a []byte
	for {
		//a, err := bufio.NewReader(out).ReadBytes('\n')
		//a, err := bufio.NewReader(out).ReadString('\n')
		a, _, err := bufio.NewReader(out).ReadLine()
		//_, err := out.Read(a)
		fmt.Print("=======JASO=======")
		fmt.Println(string(a))
		if err != nil {
			fmt.Print("1")
			fmt.Println(err)
			os.Exit(1)
		}

		_, err = bufout.WriteString(string(a))
		if err != nil {
			fmt.Print("2")
			fmt.Println(err)
			os.Exit(2)
		}
		_, err = bufout.WriteString("\n")
		if err != nil {
			fmt.Print("3")
			fmt.Println(err)
			os.Exit(3)
		}
		err = bufout.Flush()
		if err != nil {
			fmt.Print("4")
			fmt.Println(err)
			os.Exit(4)
		}
		if string(a) == "[IXAdaemon]EOD" {
			fmt.Println("manageCommunicationek [IXAdaemon]EOD jaso du")
			break
		}
	}
	c <- struct{}{}
	return

}
*/
