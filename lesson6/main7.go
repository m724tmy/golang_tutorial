package main

import (
	"fmt"
	"os"
)

//defar

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

//defer文で登録された式は、終了時にあとから登録された式から順番に評価されるシステム

func main() {
	TestDefer()

	/*
	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
	*/

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
//defer文でリソース解放処理を防ぐためにちゃんと閉ｓｓじる
	file.Write([]byte("Hello"))

}

//なぜdeferを使ったらfileが生成されるの？