package main

import "fmt"

// init

func init() {
	fmt.Println("init")
}

// initを複数に分けるメリットは通常ない
func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}
