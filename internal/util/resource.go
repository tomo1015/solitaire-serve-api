package util

import (
	"solitaire-serve-api/internal/models"
	"time"
)

const GainPerSecondPerBuilding = 1

func CollectResources(player *models.Player) {
	now := time.Now()

	for _, b := range player.Buildings {
		duration := now.Sub(b.LastCollected).Seconds()
		earned := int(duration) * b.Production

		//資源上限の計算
		warehouseLevel := player.GetWarehouseLevel()
		capacity := CollectResourceCap(warehouseLevel)

		//施設ごとに資源をそれぞれ加算する
		switch b.ResourceType {
		case 1:
			player.Resources.Wood += earned
			if player.Resources.Wood > capacity {
				player.Resources.Wood = capacity
			}
		case 2:
			player.Resources.Stone += earned
			if player.Resources.Stone > capacity {
				player.Resources.Stone = capacity
			}
		case 3:
			player.Resources.Gold += earned
			if player.Resources.Gold > capacity {
				player.Resources.Gold = capacity
			}
		}
		b.LastCollected = now
	}
}

// 資源上限計算関数
func CollectResourceCap(level int) int {
	baseCap := 1000
	increment := (level / 3) * 500
	return baseCap + increment
}
