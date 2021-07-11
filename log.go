package main

// To do
// [ ] some bug with the file permission type

import (
	"log"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

func main() {
	writeLog(INFO, "This is an information message.")
	writeLog(WARNING, "This is an warning message.")
	writeLog(ERROR, "This is an error message.")
	writeLog(FATAL, "We crashed.")
	writeLog(INFO, "You should not see this message.")

}

func writeLog(messagetype messageType, message string) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 066)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	switch messagetype {
	case INFO:
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case WARNING:
		logger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case ERROR:
		logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case FATAL:
		logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(message)
	}
}
