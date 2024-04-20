package main

import "fmt"

// 関数
// 関数を返す関数

func ReturnFunc() func() {
	return func() {
		fmt.Println(("I'm a fucntion"))
	}
}

func main() {
	f := ReturnFunc()
	f()

}
