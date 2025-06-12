package storage

import (
	"solitaire-serve-api/internal/models"
	"sync"
)

var Players = make(map[string]*models.Player) //プレイヤー情報
var Attacks []*models.Attack                  //攻撃情報
var DefensePoints []*models.DefensePoint      //防衛拠点の位置
var mu sync.Mutex

func GetPlayer(id string) *models.Player {
	mu.Lock()
	defer mu.Unlock()
	return Players[id]
}

func SavePlayer(p *models.Player) {
	mu.Lock()
	defer mu.Unlock()
	Players[p.ID] = p
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
