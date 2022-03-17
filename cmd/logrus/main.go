package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Logrus")

	//create log file and set file permissions 644
	//owner granted read/write, everyone else granted read
	file, err := os.OpenFile("cloudacademy.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.SetFormatter(&log.JSONFormatter{})

	log.Debug("debugging level information here...")
	log.Info("infomational level information here...")
	log.Warn("warning level information here...")
	log.Error("exception/error level information here...")

	fmt.Println()
}
