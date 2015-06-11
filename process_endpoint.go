package main

import (
	"bufio"
)

type ProcessEndpoint struct {
	process    *LaunchedProcess
	bufferedIn *bufio.Writer
	output     chan string
}

func NewProcessEndpoint(process *LaunchedProcess) *ProcessEndpoint {
	return &ProcessEndpoint{
		process:    process,
		bufferedIn: bufio.NewWriter(process.stdin),
		output:     make(chan string)}
}

func (pe *ProcessEndpoint) Terminate() {
}

func (pe *ProcessEndpoint) Output() chan string {
	return pe.output
}

func (pe *ProcessEndpoint) Send(msg string) bool {
	pe.bufferedIn.WriteString(msg)
	pe.bufferedIn.WriteString("\n")
	pe.bufferedIn.Flush()
	return true
}

func (pe *ProcessEndpoint) StartReading() {
	go pe.process_stdout()
}

func (pe *ProcessEndpoint) process_stdout() {
	bufin := bufio.NewReader(pe.process.stdout)
	for {
		str, err := bufin.ReadString('\n')
		if err != nil {
			break
		}
		pe.output <- trimEOL(str)

	}
	close(pe.output)
}

// trimEOL cuts unixy style \n and windowsy style \r\n suffix from the string
func trimEOL(s string) string {
	lns := len(s)
	if lns > 0 && s[lns-1] == '\n' {
		lns--
		if lns > 0 && s[lns-1] == '\r' {
			lns--
		}
	}
	return s[0:lns]
}
