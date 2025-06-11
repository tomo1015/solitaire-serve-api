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

		//施設ごとに資源をそれぞれ加算する
		switch b.ResourceType {
		case "wood":
			player.Resources.Wood += earned
		case "stone":
			player.Resources.Stone += earned
		case "gold":
			player.Resources.Gold += earned
		}

		b.LastCollected = now
	}
}
