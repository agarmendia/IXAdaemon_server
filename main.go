package main

import (
	"fmt"
)

var lProcess *LaunchedProcess

func main() {

	mainPort, ctrlPort, command, args := parseArguments()

	//Launch native program, redirecting in, out, error pipes
	lProcess = launchNative(command, args)

	interceptSignals(lProcess.cmd)

	//Launch server
	mainListener, ctrlListener := initializeServer(mainPort, ctrlPort)

	//
	go listenState(lProcess.stderr, ctrlListener)

	for {
		//Accept Client
		fmt.Println("Server listening on port: " + mainPort + "\n")
		conn, err := mainListener.AcceptTCP()

		if err != nil {
			continue
		}

		err = manageCommunication(conn, ctrlPort, command, args)
		if err != nil {
			fmt.Println(err.Error())
			lProcess = fixNative(command, args, ctrlListener)
		}
		fmt.Println("Conexion closed \n")
		conn.Close()
	}

}
