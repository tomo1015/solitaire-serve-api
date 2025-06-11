package util

import (
	"solitaire-serve-api/internal/models"
	"time"
)

const GainPerSecondPerBuilding = 1

func CollectResources(player *models.Player) {
	now := time.Now().Unix()

	if player.LastCollected == 0 {
		player.LastCollected = now
		return
	}

	elapsed := now - player.LastCollected
	if elapsed <= 0 {
		return
	}

	//施設数 * 経過秒数 * 係数
	totalGain := int(elapsed) * len(player.Buildings) * GainPerSecondPerBuilding
	player.Resources += totalGain
	player.LastCollected = now
}
