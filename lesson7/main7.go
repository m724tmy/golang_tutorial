package main

import "fmt"

//マップ
//for

func main() {
	m := map[string]int{
		"Apple": 100,
		"Banana": 200,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	//keyを使わない場合
	for _, v := range m {
		fmt.Println(v)
	}

	//valueを使わない場合
	for k, _:= range m {
		fmt.Println(k)
	}
}