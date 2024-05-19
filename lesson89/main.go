package main

import (
	"fmt"
	"regexp"
)

// regexp

// func main() {
// 	// Goの正規表現の基本 regexp.MatchString
// 	match, _ := regexp.MatchString("A", "ABC")
// 	fmt.Println(match)

// 	// Compile
// 	rel, _ := regexp.Compile("A")
// 	match = rel.MatchString("ABC")
// 	fmt.Println(match)

// 	// MustCompile
// 	re2 := regexp.MustCompile("A")
// 	match = re2.MatchString("ABC")
// 	fmt.Println(match)

// 	regexp.MustCompile("\\d")
// 	regexp.MustCompile(`\d`)
// }

func main() {
	// 正規表現のフラグ

	/*
		フラグ一覧

		i 大文字小文字を区別しない
		m マルチラインモード（^と&が文頭、文末に加えて行頭、行末にマッチ)
		s .が\nにマッチ
		u 最小マッチへの変換（X*はX*?へ、X+はX+?へ）
	*/
	re3 := regexp.MustCompile(`(?i-ms)abc`)
	match := re3.MatchString("ABC")
	fmt.Println(match)

	// 幅をもたない正規表現のパターン
	/*
		パターン一覧

		^ 文頭（mフラグが有効な場合は行頭にも）
		$ 文末（mフラグが有効な場合は行末にも）
		\A 文頭
		\z 文末
		\b ASCIIによるワード境界
		\B 非ASCIIによるワード境界
	*/

	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match)

	match = re4.MatchString("   ABC   ")
	fmt.Println(match)

	// 繰り返しのを表す正規表現

	re6 := regexp.MustCompile("a+b*")
	fmt.Println(re6.MatchString("ab"))
	fmt.Println(re6.MatchString("a"))
	fmt.Println(re6.MatchString("aaaabbbbbbbb"))
	fmt.Println(re6.MatchString("b"))

	// 正規表現の文字クラス
	re8 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re8.MatchString("Y"))

	re9 := regexp.MustCompile((`^[0-9A-Za-z_]{3}$`))
	fmt.Println(re9.MatchString("ABC"))
	fmt.Println(re9.MatchString("abcdefg"))

	re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(re10.MatchString(("ABC")))
	fmt.Println(re10.MatchString(("あ")))

	// 正規表現のグループ

	re11 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re11.MatchString("abcxyz"))
	fmt.Println(re11.MatchString("ABCXYZ"))
	fmt.Println(re11.MatchString("ABCxyz"))
	fmt.Println(re11.MatchString("ABCabc"))
}
