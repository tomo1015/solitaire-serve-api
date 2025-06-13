package battle

import (
	"math/rand/v2"
	"solitaire-serve-api/internal/models"
)

func ResolveBattle(attack *models.Attack, player *models.Player, defense *models.DefensePoint) {

	attackDamage := 0
	defenseDamage := 0

	//攻撃予約として設定された情報をもとに戦闘
	for _, attackSoldier := range attack.BattleSoldier {
		for _, defenseSoldier := range defense.Soldiers {
			dmg := calcDamage(attackSoldier, defenseSoldier) * attackSoldier.Quantity
			attackDamage += dmg
		}
	}

	if attackDamage > defenseDamage {
		//勝利したので資源を獲得
		switch defense.LocationType {
		case "forest":
			player.Resources.Wood += defense.Loot.Wood
		case "quarry":
			player.Resources.Stone += defense.Loot.Stone
		case "gold":
			player.Resources.Gold += defense.Loot.Gold
		}

		//勝利したので戦果を獲得(=リーダーボードのランキング付けに使用)
		player.Point += 100

		attack.Result = "win"
	} else {
		//敗北したのでポイントの獲得はなし
		player.Point += 0

		attack.Result = "lose"
	}

	attack.Processed = true
}

func calcDamage(attacker, defender *models.BattleSoldier) int {
	//命中判定
	if rand.Float64() > attacker.HitRate {
		return 0
	}

	// 基本ダメージ = 合計攻撃力 - 合計防御力（ダメージ0もあり）
	var totalAttackPower = attacker.Attack
	var totalDefensePower = defender.Defense

	baseDamage := totalAttackPower - totalDefensePower
	if baseDamage < 0 {
		return 0 //ダメージを与えられなかった
	}

	//クリティカル判定
	critMultiplier := 1.0
	if rand.Float64() < attacker.CritRate {
		critMultiplier = 1.5
	}

	return int(float64(baseDamage) * critMultiplier)
}
