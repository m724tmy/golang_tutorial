package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

// POST用のハンドラー
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// POST メソッド以外は拒否する
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// リクエストボディからJSONをデコードする
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// コンソールにデコードした内容を出力する
	fmt.Printf("Received user: %+v\n", user)

	// 処理結果を JSON で返す（ここでは受け取った内容をそのまま返している）
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func main() {
	//GET用のユーザー一覧エンドポイント
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{ID: 1, Name: "Alice"},
			{ID: 3, Name: "Bob"},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	})

	// POST用のユーザー作成エンドポイント
	http.HandleFunc("/user", createUserHandler)

	http.ListenAndServe(":8080", nil)
}