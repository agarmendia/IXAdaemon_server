package main

import (
	"bufio"
	"fmt"
	"io"
	//"io/ioutil"
	"net"
	"time"
)

func manageCommunication(conn net.Conn, in io.WriteCloser, out io.ReadCloser, errr io.ReadCloser) {
	//var a []byte
	for {
		//Get message from client
		message, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(message)
		if err != nil {
			break
		}
		//Send message to native program
		_, err = in.Write([]byte(message))
		if err != nil {
			break
		}

		for {
			//Get message from native program
			a, _, _ := bufio.NewReader(out).ReadLine()
			if string(a) == "bukatu da" {
				break
			}
			time.Sleep(time.Second)
			fmt.Print("erantzuna: ")
			fmt.Println(string(a))
			if err != nil {
				fmt.Println("errooorr")
				break
			}
			//Send message to client
			_, err = conn.Write([]byte(string(a) + "\n"))
			if err != nil {
				break
			}

		}
	}
}
