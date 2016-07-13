package main

import (
	"io"
	"os/exec"
)

type LaunchedProcess struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	stderr io.ReadCloser
}

func launchNative(command string, args []string) *LaunchedProcess {

	//Command for executing native program
	cmd := exec.Command(command, args...)

	//Redirect in, out and error pipes
	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	out, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	erro, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	//Launch the process
	if err = cmd.Start(); err != nil {
		panic(err)
	}

	dlog.Println("The native process is running \n")

	return &LaunchedProcess{cmd, in, out, erro}
}
