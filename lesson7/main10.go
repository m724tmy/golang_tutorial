package main

import (
	"fmt"
	"time"
)

//channel
//close

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)

	/*
	ch1 <- 1
	
	close(ch1)
	
	// ch1 <- 1

	//fmt.Println(<-ch1)

	i, ok := <-ch1
	fmt.Println(i, ok)

	i2, ok := <-ch1
	fmt.Println(i2, ok)
	// i = チャネルから受け取った値
	// ok = チャネルの状態を示す真偽値
	*/

	go reciever("1.groutin", ch1)
	go reciever("2.groutin", ch1)
	go reciever("3.groutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)

}