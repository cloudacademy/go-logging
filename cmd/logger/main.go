package main

import (
	"fmt"
	"log"
)

type blah struct {
	name  string
	count int
}

func main() {
	fmt.Println("Logger")

	blah := blah{
		name:  "CloudAcademy",
		count: 30,
	}

	var (
		name   string = "cloudacademy"
		length int    = 100
	)

	log.Println("CloudAcademy logging event...")
	log.Printf("data points: %s %d", name, length)
	log.Printf("blah struct instance -> %#v", blah)

	fmt.Println()
}
