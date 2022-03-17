package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("FileLogger")

	//create log file and set file permissions 644
	//owner granted read/write, everyone else granted read
	file, err := os.OpenFile("cloudacademy.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("CloudAcademy logging event...")
}
