package main

import (
	"fmt"
	"log"
	"os"
)

// os

func main() {
	// ファイル操作
	// Read
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)

	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

	bs2 := make([]byte, 128)

	nn, err := f.ReadAt(bs2, 10)
	fmt.Println(nn)
	fmt.Println(string(bs2))
}
