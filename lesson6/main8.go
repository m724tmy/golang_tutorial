package main

import "fmt"

// panic recover
// panic = ランタイムエラーを発生させて実行中のプログラムを強制終了させる機能

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("Start")
}