package main

import "fmt"

// 配列型

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	// 配列の更新↓
	fmt.Println(arr2[0])
	fmt.Println(arr2[2-1])

	fmt.Println("------------")
	// 値の更新
	arr2[2] = "C"
	fmt.Println(arr2)


	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)
	// 要素数が異なる場合は全く別の型とみなされるためエラー


	// 配列型　＝　要素数変更不可
	// スライス型　＝　要素数　変更可能
}