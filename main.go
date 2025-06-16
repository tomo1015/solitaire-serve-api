package main

import (
	"log"
	"net/http"
	"solitaire-serve-api/internal/db"
	"solitaire-serve-api/internal/handlers"
	"solitaire-serve-api/internal/scheduler"
	"solitaire-serve-api/storage"
)

func main() {
	//DB接続の追加
	db.Init()

	//スケジューラー起動
	go scheduler.Start()

	//防衛地点データをjsonから読み込む
	storage.LoadDefensePointFromJson("data/defense_point.json")

	//HTTPルーティング
	http.HandleFunc("/getGameToken", handlers.GetGameTokenHandler) //ゲームトークン取得
	http.HandleFunc("/player", handlers.HandlePlayer)              //プレイヤー取得
	//http.HandleFunc("/attack", handlers.HandleAttack)
	//http.HandleFunc("/leaderboard", handlers.HandleLeaderboard)  //リーダーボード
	http.HandleFunc("/build", handlers.BuildHandler)             //施設を建築する
	http.HandleFunc("/list", handlers.BuildingListHandler)       // 施設一覧
	http.HandleFunc("/upgrade", handlers.UpgradeBuildingHandler) //施設のアップグレード
	http.HandleFunc("/soldier", handlers.GetSoldiersHandler)     //兵士一覧
	http.HandleFunc("/soldier/upgrade", handlers.TrainHandler)   //兵士の訓練

	log.Println("サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
