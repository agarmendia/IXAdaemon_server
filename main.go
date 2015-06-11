package main

import (
	"fmt"
)

func main() {

	mainPort, ctrlPort, command, args := parseArguments()

	//Launch native program, redirecting in, out, error pipes
	LaunchedProcess := launchNative(command, args)

	//interceptSignals(LaunchedProcess.cmd)

	//Launch server
	mainListener, ctrlListener := initializeServer(mainPort, ctrlPort)
	pe := NewProcessEndpoint(LaunchedProcess)

	go listenState(LaunchedProcess.stderr, ctrlListener)

	for {
		//Accept Client
		fmt.Println("Server listening on port: " + mainPort + "\n")
		conn, err := mainListener.Accept()
		if err != nil {
			continue
		}
		se := newServerEndpoint(conn)

		PipeEndpoints(pe, se)

		fmt.Println("Conexion closed \n")
		conn.Close()
	}

}
