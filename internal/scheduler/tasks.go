package scheduler

import (
	"solitaire-serve-api/internal/battle"
	"solitaire-serve-api/storage"
	"time"
)

func Start() {
	ticker := time.NewTicker(10 * time.Second)
	now := time.Now().Unix()

	for range ticker.C {
		for _, p := range storage.Players {

			//資源を時間経過で追加する
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

			//兵士訓練の経過
			for _, c := range p.Soldiers {
				if c.Training && c.TrainingEnd <= now {
					c.Training = false
					//訓練が完了した通知を行う
				}
			}
		}

		//予約中のバトル処理を順番に実行
		AllBattles()
	}
}

func AllBattles() {
	for _, atk := range storage.Attacks {
		if atk.Processed {
			continue
		}

		player := storage.GetPlayer(atk.AttackerID)
		defense := storage.FindDefensePointByLocation(atk.Target.X, atk.Target.Y)

		if player == nil || defense == nil {
			continue
		}

		battle.ResolveBattle(atk, player, defense)
	}
}
