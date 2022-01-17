package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	//create log file and set file permissions 644
	//owner granted read/write, everyone else granted read
	file, err := os.OpenFile("cloudacademy.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.SetFormatter(&log.JSONFormatter{})

	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
}
