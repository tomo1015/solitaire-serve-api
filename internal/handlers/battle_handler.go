package handlers

import (
	"encoding/json"
	"net/http"
	"solitaire-serve-api/internal/models"
	"solitaire-serve-api/storage"
	"time"

	"github.com/google/uuid"
)

type DispatchedSoldier struct {
	Type     string `json:"type"`     //派遣する兵士のタイプ
	Level    int    `json:"level"`    //派遣する兵士のレベル
	Quantity int    `json:"quantity"` //派遣する兵士の数
}

type AttackRequest struct {
	AttackerID string              `json:"attacker_id"`
	DefenderID string              `json:"defender_id"`
	Soldiers   []DispatchedSoldier `json:"soldiers"` //どの兵士を派遣するか
}

func AttackRequestHandler(w http.ResponseWriter, r *http.Request) {
	var req AttackRequest
	//攻撃予約登録(予約してから30秒後に実行される)
	executeAt := time.Now().Add(30 * time.Second).Unix()

	attack := &models.Attack{
		ID:         uuid.New().String(),
		AttackerID: req.AttackerID,
		ExecuteAt:  executeAt,
		Processed:  false,
	}

	storage.Attacks = append(storage.Attacks, attack)
}

func GetBattleResultHandler(w http.ResponseWriter, r *http.Request) {
	attackID := r.URL.Query().Get("attack_id")

	for _, a := range storage.Attacks {
		if a.AttackerID == attackID {
			json.NewEncoder(w).Encode(a)
			return
		}
	}
	http.Error(w, "not found", http.StatusNotFound)
}
