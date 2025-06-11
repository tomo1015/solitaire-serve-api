package handlers

import (
	"encoding/json"
	"net/http"
	"solitaire-serve-api/internal/models"
	"solitaire-serve-api/storage"

	"github.com/google/uuid"
)

type BuildRequest struct {
	PlayerID string //プレイヤーID
	Name     string //建物名
}

func BuildHandler(w http.ResponseWriter, r *http.Request) {
	var req BuildRequest

	//リクエストのエラーチェック
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid Request", http.StatusBadRequest)
		return
	}

	//ストレージからプレイヤー情報を取得
	player := storage.GetPlayer(req.PlayerID)
	if player != nil {
		http.Error(w, "player is not found", http.StatusBadRequest)
		return
	}

	//建築する建物情報をまとめる
	building := models.Building{
		ID:       uuid.NewString(),
		Name:     req.Name,
		Level:    1,
		Position: len(player.Buildings),
	}

	//リストに追加した上で保存実施
	player.Buildings = append(player.Buildings, building)
	storage.SavePlayer(player)

	json.NewEncoder(w).Encode(building)
}
