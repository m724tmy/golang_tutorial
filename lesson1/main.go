package main

import "fmt"

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i1 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i1, s2)

	var i3 int
	var s3 string
	fmt.Println(i3, s3) 	// 出力結果：0

	// 値格納（追加）↓
	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 値の更新
	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	// 変数名　:= 値
	// 関数内でしか定義ができない
	i4 := 450
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// := での変数再定義不可
	// i4 := 500
	// fmt.Println(i4)

	// i4 := "Hello"
	// fmt.Println(i4)
	// 最初に(i4 : 450(int))定義された型として定義されるため、文字列出力不可


	outer()
	var s5 string = "Not Use"

	fmt.Println(s5)
}