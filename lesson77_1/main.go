package main

import (
	"fmt"
	"os"
)

// os

func main() {
	// Exit
	/*
		os.Exit(1)
		fmt.Println("Start")
	*/

	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)
}
