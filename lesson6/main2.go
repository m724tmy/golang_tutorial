package main

import (
	"fmt"
	"strconv"
)

//if
//条件分岐
//エラーハンドリング

func main () {
	var s string = "tt"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n",i)
}