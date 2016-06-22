package main

import (
	"fmt"
	"net"
)

func fixNative(command string, args []string, ctrlListener *net.TCPListener) *LaunchedProcess {
	fmt.Println("Fixing Native")
	lp := launchNative(command, args)
	fmt.Println("Native launched")
	listenState(lProcess.stderr, ctrlListener)
	return &LaunchedProcess{lp.cmd, lp.stdin, lp.stdout, lp.stderr}
}
