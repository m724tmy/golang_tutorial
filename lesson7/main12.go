package main

import "fmt"

//channel
//select

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)
	//送受信するバッファ付きチャネル

	/*
	ch2 <- "A"
	ch1 <- 1
	ch2 <- "B"
	ch1 <- 2
	*/

	/*
	v1 := <-ch1
	v2 := <-ch2
	fmt.Println(v1)
	fmt.Println(v2)
	*/

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}
	//複数のチャネル操作の中から実行可能なケースをランダムに選定。

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	//ch3: 最初の入力を受け取るためのチャネル。
	//ch4: 中間結果を保存するためのチャネル。
	//ch5: 最終結果を保存するためのチャネル。


	//reciever
	go func() {
		for {
			i := <- ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <- ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch3 <- n:
			// nをch3に送信
			n++
		case i3 := <- ch5:
			//ch5から値を受信し,i3に格納
			fmt.Println("recieved", i3)
		default:
			if n > 100 {
				break L //名前付きforループでループから抜け出す
			}
		}
		/*
		if n > 100 {
			break
		}
		*/
	}
}