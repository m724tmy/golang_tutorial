package main

import (
	"fmt"
	"time"
)

//go goroutin
// 並行処理

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub()
// 本来の処理としては、subの呼び出しが終了されてからmain（Main Loop）を呼び出す形となるが、
//延々と繰り返されるため、mainは呼び出されない。そこで、
//go sub として頭に"go "をつけるだけで並行処理の実装が可能となる。

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}