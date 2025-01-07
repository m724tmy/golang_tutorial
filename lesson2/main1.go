package main

// int型
import "fmt"

func main() {
    var i int = 100

    var i2 int64 = 200

    fmt.Println(i + 50)
    
    fmt.Printf("%T\n", i2)

    // Printf + "%T\n"を使うことで型を調べることができる
    fmt.Printf("%T\n", int32(i2))
    // 型は変換ができる
    fmt.Println(i + int(i2))
}