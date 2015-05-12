package main

import (
	//"bufio"
	//"fmt"
	"io"
	//"os"
	"os/exec"
	//"strings"
)

func konexioaEgin() (io.WriteCloser, io.ReadCloser, io.ReadCloser) {
	/*
		fmt.Println("sartu komandoa")
		var input string
		_, err := fmt.Fscanf(os.Stdin, "%s", &input)
		command := strings.Split(input, " ")

		fmt.Println(command[0])
		fmt.Println(command[1])
		cmd := exec.Command(command[0], command[1])
	*/
	cmd := exec.Command("java", "alderantziz")

	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	//defer in.Close()

	out, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}

	//defer out.Close()

	erro, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	//defer erro.Close()

	// Start the process
	if err = cmd.Start(); err != nil {
		panic(err)
	}

	return in, out, erro
}
