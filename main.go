package main

import (
	"fmt"
)

func main() {

	mainPort, ctrlPort, command := parseArguments()

	//Launch native program, redirecting in, out, error pipes
	in, out, errr := launchNative(command)
	//Launch server
	mainListener, _ := initializeServer(mainPort, ctrlPort)

	//go listenState(errr, ctrlPort)

	for {
		//Accept Client
		fmt.Println("Server listening on port: " + mainPort + "\n")
		conn, err := mainListener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Conexion is opened \n")
		//Manage communication between client and native program
		manageCommunication(conn, in, out, errr)
		fmt.Println("Conexion is closed \n")
		conn.Close()
	}

}
