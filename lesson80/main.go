package main

import (
	"fmt"
	"math/rand"
	"time"
)

// rand

func main() {
	rand.Seed(256)

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())

	// 現在の時刻をシードに使った疑似乱数の生成
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	// 0~99の間の擬似乱数
	fmt.Println(rand.Intn(100))
	// int型の擬似乱数
	fmt.Println(rand.Int())
	// int32型の擬似乱数
	fmt.Println(rand.Int31())
	// int64型の擬似乱数
	fmt.Println(rand.Int63())
	// uint32型の擬似乱数
	fmt.Println(rand.Uint32)
}
