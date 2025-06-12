package storage

import (
	"solitaire-serve-api/internal/db"
	"solitaire-serve-api/internal/models"
)

var Players = make(map[string]*models.Player) //プレイヤー情報
var Attacks []*models.Attack                  //攻撃情報
var DefensePoints []*models.DefensePoint      //防衛拠点の位置

// プレイヤーをDBから取得
func GetPlayer(id string) *models.Player {
	var player models.Player
	result := db.DB.First(&player, "id = ?", id)
	if result.Error != nil {
		return nil //見つからない場合
	}
	return &player
}

// プレイヤー情報をDBに保存
func SavePlayer(player *models.Player) {
	db.DB.Save(player)
}

// 指定された座標に存在する防衛拠点を検索する
func FindDefensePointByLocation(x, y int) *models.DefensePoint {
	for _, dp := range DefensePoints {
		if dp.Location.X == x && dp.Location.Y == y {
			return dp
		}
	}

	return nil
}
