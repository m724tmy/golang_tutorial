package main

import "fmt"

// init

func init() {
	fmt.Println("init")
}
// 本来はmainが先に読み込まれる構文だが、init関数を使うことでinitが先に読み込まれる。

// GPT解説
// init関数は、main関数よりも先に自動的に実行される特別な関数。
// 初期化処理を行うために使われる。
// main関数より前に実行されるため、初期設定やリソースの準備に適している。

func main() {
	fmt.Println("Main")
}

