package main

import (
	"log"
	"os"
)

// os

func main() {
	// ファイル操作
	// Open
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	// deferでクローズする。確実にファイルがクローズされる。
	// リソース破棄のための処理漏れや、分散したりといった問題を防ぐ
	defer f.Close()
}
