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
		player.Resources += earned
		b.LastCollected = now
	}
}
