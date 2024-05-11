package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// ファイル操作
	// OpenFile
	// O_RDONLY 読み込み専用
	// O_WRONLY 書き込み専用
	// O_RDWR 読み書き専用
	// O_APPEND ファイル末尾に追加
	// O_CREATE ファイルがなければ作成
	// O_TRUNC 可能であればファイルの内容をオープン時に空にする
	// f, err := os.OpenFile("foo.txt", os.O_RDONLY, 0666)
	f, err := os.OpenFile("foo.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}
