package main

import "fmt"

// byte(unit8)型

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	fmt.Printf("%T\n", byteA)

	// byte(unit8) => 文字列
	fmt.Println(string(byteA))

	// 文字列
	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))
}