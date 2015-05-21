package main

import (
	"fmt"
)

func main() {

	mainPort, ctrlPort, command := parseArguments()

	//Launch native program, redirecting in, out, error pipes
	in, out, errr := launchNative(command)
	//Launch server
	mainListener, ctrlListener := initializeServer(mainPort, ctrlPort)

	go listenState(errr, ctrlListener)

	for {
		//Accept Client
		fmt.Println("Server listening on port: " + mainPort + "\n")
		conn, err := mainListener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Conexion opened \n")
		//Manage communication between client and native program
		manageCommunication(conn, in, out, errr)
		fmt.Println("Conexion closed \n")
		conn.Close()
	}

}
