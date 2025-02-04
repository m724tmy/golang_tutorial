package main

import (
	"fmt"
	"net/http"
)


func helloHandler(w http.ResponseWriter, r *http.Request) {
	// クライアントに"Hello, World"を返す"
	fmt.Println(w, "Hello, World!")
}

func main() {
	// URLパス"/"に対して helloHandlerを登録
	http.HandleFunc("/", helloHandler)

	// サーバーをポート 8080 で起動し、リクエストを待ち受ける
	http.ListenAndServe(":8080", nil)
}