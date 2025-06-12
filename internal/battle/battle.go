package battle

import (
	"solitaire-serve-api/internal/models"
)

func ResolveBattle(atk *models.Attack, player *models.Player, defense *models.DefensePoint) {
	attackPower := calcTotalPower(atk.Soldiers)
	defensePower := calcTotalPower(defense.Soldiers)

	if attackPower > defensePower {
		//勝利したので資源を獲得
		switch defense.LocationType {
		case "forest":
			player.Resources.Wood += defense.Loot.Wood
		case "quarry":
			player.Resources.Stone += defense.Loot.Stone
		case "gold":
			player.Resources.Gold += defense.Loot.Gold
		}
		atk.Result = "win"
	} else {
		atk.Result = "lose"
	}

	atk.Processed = true
}

// 兵士の戦力計算関数
func calcTotalPower(soldiers []*models.Soldier) int {
	power := 0
	for _, s := range soldiers {
		power += s.Quantity * s.Level //戦力 = 敵 × レベル
	}
	return power
}
