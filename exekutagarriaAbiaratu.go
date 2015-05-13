package main

import (
	"fmt"
	"io"
	"os/exec"
)

func exekutagarriaAbiaratu(komandoa []string) (io.WriteCloser, io.ReadCloser, io.ReadCloser) {

	//exekutatuko den komandoa
	cmd := exec.Command(komandoa[0], komandoa[1])

	//Sarrera, irteera eta errore pipeak berbideratu
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

	//Prozesua martxan jarri
	if err = cmd.Start(); err != nil {
		panic(err)
	}

	fmt.Println("Prozesua martxan da \n")

	return in, out, erro
}
