package handlers

import (
	"encoding/json"
	"net/http"
	"solitaire-serve-api/internal/models"
	"solitaire-serve-api/storage"
	"time"

	"github.com/google/uuid"
)

type TrainRequest struct {
	PlayerID string `json:"player_id"` //プレイヤーID
	Type     string `json:"type"`      //兵士タイプ
	Quantity int    `json:"quantity"`  //訓練数
}

// 兵士訓練実行API
func TrainHandler(w http.ResponseWriter, r *http.Request) {
	var req TrainRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid Request", http.StatusBadRequest)
		return
	}

	//ストレージからプレイヤー情報を取得
	player := storage.GetPlayer(req.PlayerID)
	if player == nil {
		http.Error(w, "player is not found", http.StatusBadRequest)
		return
	}

	//コストチェックや訓練中判定を追加

	now := time.Now().Unix()
	trainDuration := int64(60) //60秒の訓練時間

	soldier := &models.Soldier{
		ID:          uuid.New().String(),
		Type:        req.Type,
		Level:       1,
		Quantity:    req.Quantity,
		Training:    true,
		TrainingEnd: now + trainDuration,
	}

	player.Soldiers = append(player.Soldiers, soldier)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("training started"))
}

// 兵士の一覧取得API
func GetSoldiersHandler(w http.ResponseWriter, r *http.Request) {
	playerID := r.URL.Query().Get("player_id")
	player := storage.GetPlayer(playerID)

	if player == nil {
		http.NotFound(w, r)
		return
	}
}
