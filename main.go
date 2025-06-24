package main

import (
	"log"
	"net/http"
	"solitaire-serve-api/docs"
	"solitaire-serve-api/internal/db"
	"solitaire-serve-api/internal/handlers"
	"solitaire-serve-api/internal/scheduler"
	"solitaire-serve-api/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	//DB接続の追加
	db.Init()

	//スケジューラー起動
	go scheduler.Start()

	//防衛地点データをjsonから読み込む
	storage.LoadDefensePointFromJson("data/defense_point.json")
	storage.LoadFacilityFromJson("data/facility.json")

	//Ginのルーターを取得
	r := gin.Default()

	// docsパッケージ内の SwaggerInfo を使う（使わなければunusedになる）
	docs.SwaggerInfo.BasePath = "/"

	// SwaggerUIのルーティング
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//エンドポイント
	r.POST("/getGameToken", handlers.GetGameTokenHandler)
	r.POST("/facility/create", handlers.FacilityHandler)
	r.POST("/facility/list", handlers.FacilityListHandler)
	r.POST("/facility/upgrade", handlers.UpgradeFacilityHandler)
	r.POST("/soldier/list", handlers.GetSoldiersHandler)
	r.POST("/soldier/training", handlers.TrainHandler)

	//HTTPルーティング
	// http.HandleFunc("/player", handlers.HandlePlayer)              //プレイヤー取得
	// //http.HandleFunc("/attack", handlers.HandleAttack)
	// //http.HandleFunc("/leaderboard", handlers.HandleLeaderboard)  //リーダーボード

	log.Println("サーバー起動: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
