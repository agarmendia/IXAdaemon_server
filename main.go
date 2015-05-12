package main

import (
	//"bufio"
	"fmt"
)

func main() {

	in, out, errr := konexioaEgin()

	Listener := zerbitzaria()

	fmt.Println("Prozesua martxan da \n")

	for {
		conn, err := Listener.Accept()
		if err != nil {
			continue
		}
		bezeroaKudeatu(conn, in, out, errr)
		// will listen for message to process ending in newline (\n)
		conn.Close()
	}

}
