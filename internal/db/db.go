package db

import (
	"log"
	"solitaire-serve-api/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("game.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}

	log.Println("DB接続成功")

	//テーブル作成の実行
	DB.AutoMigrate(&models.Player{}, &models.Building{}, &models.Attack{}, &models.DefensePoint{}, &models.Soldier{})
}
