package main

import (
	"fmt"
	"log"
	"os"
)

var errorlog *os.File
var logger *log.Logger

func initializeLogger(logfile string) *log.Logger {
	if logfile != "" {
		errorlog, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
			os.Exit(1)
		}
		logger = log.New(errorlog, "applog: ", log.Lshortfile|log.LstdFlags)
	} else {
		logger = log.New(os.Stdout, "applog: ", log.Lshortfile|log.LstdFlags)
	}

	return logger
}
