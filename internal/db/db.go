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
	// DSNの先頭に "file:" をつけるとmodernc.org/sqliteドライバで認識されやすいです
	DB, err = gorm.Open(sqlite.Open("file:game.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB接続失敗: %v", err)
	}

	log.Println("DB接続成功")

	DB.AutoMigrate(
		&models.Session{},
		&models.DefensePoint{},
		&models.BattleSoldier{},
		&models.Resources{},
		&models.Building{})
}
