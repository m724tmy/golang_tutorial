package main

import "fmt"

// 関数
// ジェネレーターの実装

func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
		// 実行されるたびに参照される
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
}