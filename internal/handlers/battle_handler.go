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
	AttackerID  string          `json:"attacker_id"`
	TargetPoint models.WorldMap `json:"location"`
	SoldierID   int             `json:"soldier_id"`
	Quantity    int             `json:"quantity"`
}

func AttackRequestHandler(w http.ResponseWriter, r *http.Request) {
	var req AttackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//プレイヤードキュメントから兵士を取得
	player := storage.GetPlayer(req.AttackerID)
	soldier := player.FindSoldier(req.SoldierID)
	if soldier == nil || soldier.Quantity < req.Quantity {
		http.Error(w, "Not Enough soldiers", http.StatusBadRequest)
		return
	}

	//予約の段階で使用する分を差し引く
	soldier.Quantity -= req.Quantity
	storage.SavePlayer(player)

	//BattleSoldierに複写
	battleSoldier := &models.BattleSoldier{
		Type:     soldier.Type,
		Level:    soldier.Level,
		Quantity: req.Quantity,
		Attack:   soldier.Level * 10, //例（本来はレベルごとの曲線にしたい）
		Defense:  soldier.Level * 5,  //例
		CritRate: 0.07,
		HitRate:  0.9, //
	}

	//攻撃予約登録(予約してから30秒後に実行される)
	executeAt := time.Now().Add(30 * time.Second).Unix()

	attack := &models.Attack{
		ID:            uuid.New().String(),
		AttackerID:    req.AttackerID,
		Target:        req.TargetPoint,
		BattleSoldier: []*models.BattleSoldier{battleSoldier},
		ExecuteAt:     executeAt,
		Processed:     false,
	}

	//DBに保存
	storage.Attacks = append(storage.Attacks, attack)

	w.WriteHeader(http.StatusCreated)
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
