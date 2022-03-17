package main

import (
	"fmt"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	//create log file and set file permissions 644
	//owner granted read/write, everyone else granted read
	file, err := os.OpenFile("cloudacademy.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	InfoLogger.Println("CloudAcademy app started...")
}

type blah struct {
	name  string
	count int
}

func main() {
	fmt.Println("CustomLogger")

	blah := blah{
		name:  "CloudAcademy",
		count: 30,
	}

	InfoLogger.Println("blah message")

	WarningLogger.Println("warning blah blah")
	WarningLogger.Printf("blah struct instance -> %#v", blah)

	ErrorLogger.Println("major blah blah detected!!")
	fmt.Println()
}
