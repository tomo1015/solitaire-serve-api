package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserID     int    `json:"user_id"`     //ユーザーID（一意のもの）
	ShardKey   int    `json:"shard_key"`   //シャードキー
	PlatFormID int    `json:"platform_id"` //プラットフォームID
	GameToken  string `json:"game_token"`  //ゲームトークン
}
