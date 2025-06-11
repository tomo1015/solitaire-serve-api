package scheduler

import (
	"solitaire-serve-api/storage"
	"time"
)

func Start() {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		for _, p := range storage.Players {
			p.Resources += 10 // 資源の自動増加
		}
	}
}
