package main

import "fmt"

//定数
// 定数は値の変更があとからできない
// 定数はグローバルに書くことが多い

const Pi = 3.14

const (
	URL = "http:;...sfsa"
	SiteName = "test"
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	// Pi = 3
	// fmt.Println(Pi)
	// 値の再変更不可

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(c0, c1, c2)
}