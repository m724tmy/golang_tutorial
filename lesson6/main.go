package main

import "fmt"

//if
//条件分岐

func main() {
	a := 1
	if a == 2{
		fmt.Println("two")
	} else if a == 1{
		fmt.Println("one")
	} else {
		fmt.Println("I dont know")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}
}