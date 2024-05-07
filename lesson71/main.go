package main

import "fmt"

// interface
// fmt.Stringer

/*
// このように定義されている
type Stringer interface {
	Stringer() string
}
*/

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "ABC"}
	fmt.Println(p)
}
