package main

import (
	"fmt"
	"io"
	"os/exec"
)

func launchNative(command []string) (io.WriteCloser, io.ReadCloser, io.ReadCloser) {

	//Command for executing native program
	cmd := exec.Command(command[0], command[1])

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

	fmt.Println("The native process is running \n")

	return in, out, erro
}
