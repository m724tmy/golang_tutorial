package main

import "fmt"

func main() {
	var s string = "Hello World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var si string = "3000"
	fmt.Println(si)


	// 改行出力
	fmt.Printf("%T\n", si)

	fmt.Println(`test
	test 
		testunko`)


	// ダブルクォーテーションで"を出力する
	fmt.Println("\"")

	// Hello WorldのHを出力期待したいが、下記は72と出力される。
	// 文字列はバイト配列の集まりである。
	fmt.Println(s[0])

	// 文字列1文字（H）を出力する場合
	fmt.Println(string(s[0]))

}
