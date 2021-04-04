package main

import (
	"log"
	"os"
)

func main() {
	// Create, Append, WriteOnly logs file
	logsFile, err := os.OpenFile("logs.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)

	if err != nil {
		// Fatal makes program crash
		log.Fatal(err)
	}

	log.SetOutput(logsFile)
	log.Println("Hello there")
	log.Println("Output here")
}