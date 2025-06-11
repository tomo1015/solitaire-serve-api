package scheduler

import (
	"solitaire-serve-api/storage"
	"time"
)

func Start() {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		for _, p := range storage.Players {
			for _, b := range p.Buildings {
				switch b.ResourceType {
				case "wood":
					p.Resources.Wood += b.Production
				case "stone":
					p.Resources.Stone += b.Production
				case "gold":
					p.Resources.Gold += b.Production
				}
			}
		}
	}
}
