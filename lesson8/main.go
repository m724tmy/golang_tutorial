package main

import "fmt"

type User struct {
	Name string
	Age int
	//X, Y int
	// fieldをまとめて定義可能
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3:= User{Name: "user3" , Age:30} //初期値を持たせて定義可能
	fmt.Println(user3)

	user4 := User{"user4", 40} //user3同最初からフィールドに値代入可能
	fmt.Println(user4)

	// user5 := User{30, "user5"} 
	// fmt.Println(user5)
	//NG fieldの順番に値を代入する必要があるためエラー

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)
	//user7, 8についてはポインタ型で宣言可能

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)
}