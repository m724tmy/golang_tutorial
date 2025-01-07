package main

import "fmt"

// 関数
// クロージャー

func Latter() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func main() {
	f := Latter()
	fmt.Println(f("Hello"))
	fmt.Println(f("tn"))
	fmt.Println(f("eeeo"))
	fmt.Println(f("ewro"))
}

// 理屈はわからんが、最後が出力されない。
// なあぜ・？