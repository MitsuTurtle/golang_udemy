package foo

const (
	Max = 100
	// 先頭の文字が小文字だと外部のパッケージから参照できない
	min = 1
)

func ReturnMIN() int {
	return min
}
