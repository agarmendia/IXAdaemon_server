package main

import (
	"net"
)

func fixNative(command string, args []string, ctrlListener *net.TCPListener, currentDoc []string) *LaunchedProcess {
	dlog.Println("Fixing Native")
	dlog.Println(currentDoc)
	lp := launchNative(command, args)

	go listenState(lProcess.stderr, ctrlListener)
	return &LaunchedProcess{lp.cmd, lp.stdin, lp.stdout, lp.stderr}
}
