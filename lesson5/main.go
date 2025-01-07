package main

import "fmt"

// 関数

func Plus(x int, y int) int {
// func Plus(x, y int) int  といった書き方もOK
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)
	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3)
	// i2, _ := Div(9, 3)
	// fmt.Println(i2) という書き方もOK、　_をつかうことで第２引数を省けれれる

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()
}