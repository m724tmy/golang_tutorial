package main

import "fmt"

// 浮動小数点型

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)
	// 小数点の型　＝　float64

	fl := 3.2
	// 暗黙的(:=)な変数定義は自動でfloat64となる
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)
	// float32は基本使わない
	fmt.Printf("%T\n", float64(fl32))
	// 型変換 float32 => float64
}