package main

import "fmt"

type User struct {
	Name string
	Age int
	//X, Y int
	//Userは、ユーザーの名前（Name）と年齢(Age)を持つデーター型
}

type Users []*User
//Usersは、*User（User構造体へのポインタ）のスライス型を定義。

/*
 type Users struct {
	Users []*Users
 }
*/

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	//4つのUser構造体を作成し、それぞれに名前と年齢を設定。

	users := Users{}

	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)

	//Users型のスライスusersを作成し、appendでUser構造体のポインタを追加。
	//ポインタを使うことで、スライス内の要素を変更すれば元のUser構造体も変更
	//コピーではなく参照を操作するため、メモリ効率が良い。


	for _, u := range users {
		fmt.Println(*u)
	}

	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)

	for _, u := range users2 {
		fmt.Println(*u)
	}
}



