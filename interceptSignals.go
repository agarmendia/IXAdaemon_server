package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func interceptSignals(cmd *exec.Cmd) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		switch sig {
		case syscall.SIGTERM:
			dlog.Println("Server hard close [SIGTERM]")
			cmd.Process.Kill()
			os.Exit(1)
		}
	}()
}
