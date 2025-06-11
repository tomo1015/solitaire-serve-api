package main

import (
	"log"
	"net/http"
	"solitaire-serve-api/internal/handlers"
	"solitaire-serve-api/internal/scheduler"
)

func main() {
	//スケジューラー起動
	go scheduler.Start()

	//HTTPルーティング
	http.HandleFunc("/player", handlers.HandlePlayer) //プレイヤー
	//http.HandleFunc("/attack", handlers.HandleAttack)
	http.HandleFunc("/leaderboard", handlers.HandleLeaderboard) //リーダーボード
	http.HandleFunc("/build", handlers.BuildHandler)            //施設を建築する

	log.Println("サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
