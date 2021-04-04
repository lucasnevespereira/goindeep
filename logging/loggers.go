package main

import (
	"log"
	"os"
)

var (
	ErrorLogger *log.Logger
	WarningLogger *log.Logger
	InfoLogger *log.Logger
)

func initLoggers() {
	// Create, Append, WriteOnly logs file
	logsFile, err := os.OpenFile("logs.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)

	if err != nil {
		// Fatal makes program crash
		log.Fatal(err)
	}

	// Custom loggers with prefix and (date + time + concerned file)
	ErrorLogger = log.New(logsFile, "[ERROR] ", log.Ldate | log.Ltime | log.Lshortfile)
	WarningLogger = log.New(logsFile, "[WARNING] ", log.Ldate | log.Ltime | log.Lshortfile)
	InfoLogger = log.New(logsFile, "[INFO] ", log.Ldate | log.Ltime | log.Lshortfile)
}
