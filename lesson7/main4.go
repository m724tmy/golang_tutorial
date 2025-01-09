package main

import "fmt"

// スライス
// for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// for _, v := range sl {
	// 	fmt.Println(v)
	// }

	// for i := range sl {
	// 	fmt.Println(i, sl[i])
	// }

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}