package main

import (
	"github.com/m724tmy/go-rest-api/db"
	"github.com/m724tmy/go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() //DB接続
	r := gin.Default()
	routes.SetupRoutes(r) // ルーティングの設定
	r.Run(":8080") // サーバーを起動（デフォルトは `localhost:8080`）
}
