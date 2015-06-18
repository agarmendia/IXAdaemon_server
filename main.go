package main

import (
	"fmt"
)

func main() {

	mainPort, ctrlPort, command, args := parseArguments()

	//Launch native program, redirecting in, out, error pipes
	LaunchedProcess := launchNative(command, args)

	interceptSignals(LaunchedProcess.cmd)

	//Launch server
	mainListener, ctrlListener := initializeServer(mainPort, ctrlPort)

	go listenState(LaunchedProcess.stderr, ctrlListener)

	for {
		//Accept Client
		fmt.Println("Server listening on port: " + mainPort + "\n")
		conn, err := mainListener.AcceptTCP()
		if err != nil {
			continue
		}
		manageCommunication(conn, LaunchedProcess.stdin, LaunchedProcess.stdout, LaunchedProcess.stderr)
		fmt.Println("Conexion closed \n")
		conn.Close()
	}

}
