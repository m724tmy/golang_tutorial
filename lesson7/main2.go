package main

import "fmt"

//スライス
// append make len cap

func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)
	//配列複数追加可能

	sl2 := make([]int, 5)
	fmt.Println(sl2)
	fmt.Println(len(sl2)) //len = 長さ
	fmt.Println(cap(sl2)) //cap = 容量
	// make を使って初期サイズと容量を指定したスライスを作成している。
	// len と cap を用いて長さと容量を確認している。

	sl3 := make([]int, 5, 10)
	// 長さ 5、容量 10 のスライスを作成。
	fmt.Println(len(sl3))
	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)

	fmt.Println(len(sl3))

	fmt.Println(cap(sl3))
	// 初期容量を超える要素を追加した場合、容量が２倍、またはそれ以上に確保し直す。
	// メモリ再確保が頻繁に発生するとパフォーマンス面で影響される
}