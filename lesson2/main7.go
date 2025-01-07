package main

import "fmt"

// interface & nil

func main() {
	var x interface{}
	fmt.Println(x)

	// interfaceで、データー型を異なる型で置き換え可能
	x = 1
	fmt.Println(x)
	x = 4.56
	fmt.Println(x)
	x = "hello"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	// ただし、データー特有の計算などはできない
	x = 2
	fmt.Println(x + 3)
	// インターフェイスとの計算は不可
	
}