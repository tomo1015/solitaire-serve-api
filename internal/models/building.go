package models

import "time"

type Building struct {
	ID            string    `json:"id"`            //一意なID（建設時に発行）
	BuildingID    int       `json:"facility_id"`   //施設ID
	Level         int       `json:"level"`         //施設レベル（初期1)
	Position      int       `json:"position"`      //建物の位置
	Production    int       `json:"production"`    //施設ごとの生産量
	ResourceType  string    `json:"resource_type"` //施設ごとに生産できる資源タイプ
	LastCollected time.Time `json:"last_collected"`
}

// 建設/アップグレード時のコスト計算
func (b *Building) UpgradeCost() int {
	cost := 0
	if b.Level >= 3 {
		cost = 100 * b.Level
	}
	return cost
}
