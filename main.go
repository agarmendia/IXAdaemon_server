package main

import (
	"fmt"
)

var lProcess *LaunchedProcess

func main() {

	//Parse arguments from command line
	mainPort, ctrlPort, command, args := parseArguments()

	//Launch native program, redirecting in, out, error pipes
	lProcess = launchNative(command, args)

	//Intercpets CTR+C and closes the native in that case
	interceptSignals(lProcess.cmd)

	//Launch server
	mainListener, ctrlListener := initializeServer(mainPort, ctrlPort)

	//Get status of the native for IXAdaemon_ctrl
	go listenState(lProcess.stderr, ctrlListener)

	for {
		//Accept Client
		fmt.Println("Server listening on port: " + mainPort + "\n")
		conn, err := mainListener.AcceptTCP()

		if err != nil {
			continue
		}

		//Manage Client <-> Native communication
		err = manageCommunication(conn, ctrlPort, command, args)

		if err != nil {
			fmt.Println(err.Error())

			//In case of communication goes wrong, relaunch Native
			lProcess = fixNative(command, args, ctrlListener)
		}
		fmt.Println("Conexion closed \n")
		conn.Close()
	}

}
