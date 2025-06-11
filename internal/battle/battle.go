package battle

import (
	"solitaire-serve-api/storage"
)

func Attack(attackerID, defenderID string) string {
	attacker := storage.GetPlayer(attackerID)
	defender := storage.GetPlayer(defenderID)

	if attacker == nil || defender == nil {
		return "invalid"
	}

	if attacker.Soldiers > defender.Soldiers {
		attacker.Resources += 50
		defender.Resources -= 30
		return "win"
	}

	return "lose"
}
