package main

import (
	"fmt"
	"strconv"
)

// 型変換

func main() {
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fl64 = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3 = %T\n", i3)
	// fmt.Println(i3)
	// 小数点以下は切り捨てられる。

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	// 文字列から数値への型変換
	i, _ := strconv.Atoi(s)
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	// 数値型から文字列への変換
	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	// 文字列〜バイト配列変換


	fmt.Println("^--------^")

	var h string ="Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)
}