package main

import (
	"log"
	"os"
)

// os

func main() {
	// log.Fatal
	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err)
	}
}
