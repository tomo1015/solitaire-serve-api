package models

type Session struct {
	UserID     int    `gorm:"primaryKey;autoIncrement"` //ユーザーID（一意のもの）
	ShardKey   int    `json:"shard_key"`                //シャードキー
	PlatFormID int    `json:"platform_id"`              //プラットフォームID
	GameToken  string `json:"game_token"`               //ゲームトークン
}
