package logging

import (
	"log"
	"os"
)

const path string = "server.log"

func Log(message string) {
	// If the file does not exist, create/append to it
	// If the file exists, append to it
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	} else {
		log.SetOutput(file)
		log.Println(message)
	}

	file.Close()
}
