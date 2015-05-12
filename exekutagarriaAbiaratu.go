package main

import (
	"io"
	"os/exec"
)

func exekutagarriaAbiaratu() (io.WriteCloser, io.ReadCloser, io.ReadCloser) {

	//exekutatuko den komandoa
	cmd := exec.Command("java", "alderantziz")

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

	return in, out, erro
}
