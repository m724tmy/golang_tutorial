package main

import "fmt"

// 関数
// 関数を返す関数

func ReturnFunc() func() {
	return func() {
		fmt.Println("Im a function")
	}
}

func main() {
	f := ReturnFunc()
	f()
}