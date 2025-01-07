package main

import "fmt"

// for
//繰り返し処理

func main() {
	/*
	i := 0
	for {
		i ++
		if i == 3 {
			break
		}
		fmt.Println("Loop")
	}
	*/

	/*
	point := 0
	for point <= 10 {
		fmt.Println(point)
		point++
	}
	*/

	/*
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
	*/

	/*
	//配列 ＝ 要素を格納しているデーター型
	//len :要素数を取得できる
	arr := [3]int{1, 2, 3}
	for i:=0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	*/

	//範囲式for

	/*
	arr := [3]int{1,2,3}
	for i, v:= range arr {
		fmt.Println(i, v)
	}
	*/

	/*
	sl := []string{"Python", "PHP", "GO"}
	for i, v := range sl {
		fmt.Println(i, v)
	}
	*/

	//map, key valueの形式でデーターを格納する
	m := map[string]int{"apple": 100, "banana":200}

	for k, i := range m {
		fmt.Println(k, i)
	}
}