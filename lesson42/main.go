package main

import "fmt"

// panic recover
// この方式でのエラーハンドリングは非推奨なので使用場面は少ない

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("Start")
}
