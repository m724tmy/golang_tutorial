package main

import "fmt"

//channel
//複数のゴルーチン間でのデータの受け渡しをする為に設計されたデータ構造。
//宣言、操作

func main() {
	var ch1 chan int

	//受信専用
	// var ch2 <-chan int

	//送信専用
	// var ch3 chan<- int

	//make関数を使うことでチャネルとして機能を持たせる。
	ch1 = make(chan int)

	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	i := <-ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))

	i2 := <- ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))
	// <- 矢印を使ってデータの送受信が可能

	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))

	ch3 <- 1
	fmt.Println("len", len(ch3))
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6

}