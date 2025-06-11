package handlers

import (
	"encoding/json"
	"net/http"
	"solitaire-serve-api/internal/models"
	"solitaire-serve-api/storage"
	"time"

	"github.com/google/uuid"
)

type BuildRequest struct {
	PlayerID string `json:"player_id"` //プレイヤーID
	Name     string `json:"name"`      //建物名
}

type UpgradeRequest struct {
	PlayerID   string `json:"player_id"`   //プレイヤーID
	BuildingID string `json:"building_id"` //建物ID
}

// 施設を建築する
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
	if player == nil {
		http.Error(w, "player is not found", http.StatusBadRequest)
		return
	}

	//初期レベルの建設コスト計算
	tmpBuilding := models.Building{Name: req.Name, Level: 1}
	cost := tmpBuilding.UpgradeCost()
	//コストの消費
	switch tmpBuilding.ResourceType {
	case "wood":
		if player.Resources.Wood < cost {
			http.Error(w, "resources not enough", http.StatusBadRequest)
			return
		}

		player.Resources.Wood -= cost

	case "stone":
		if player.Resources.Stone < cost {
			http.Error(w, "resources not enough", http.StatusBadRequest)
			return
		}
		player.Resources.Stone -= cost

	case "gold":
		if player.Resources.Gold < cost {
			http.Error(w, "resources not enough", http.StatusBadRequest)
			return
		}

		player.Resources.Gold -= cost
	}

	//建築時に施設ごとの生産量を設定
	production := 1
	resource_type := "wood"
	switch req.Name {
	case "木材工場":
		production = 5
		resource_type = "wood"
	case "石材工場":
		production = 3
		resource_type = "stone"
	case "金鉱山":
		production = 2
		resource_type = "gold"
	}

	//建築する建物情報をまとめる
	building := models.Building{
		ID:            uuid.NewString(),
		Name:          req.Name,
		Level:         1,
		Position:      len(player.Buildings),
		Production:    production,
		ResourceType:  resource_type,
		LastCollected: time.Now(),
	}

	//リストに追加した上で保存実施
	player.Buildings = append(player.Buildings, building)
	storage.SavePlayer(player)

	json.NewEncoder(w).Encode(building)
}

// 施設一覧
func BuildingListHandler(w http.ResponseWriter, r *http.Request) {
	playerId := r.URL.Query().Get(("player_id"))
	if playerId == "" {
		http.Error(w, "player id is missing", http.StatusBadRequest)
		return
	}

	player := storage.GetPlayer(playerId)
	if player == nil {
		http.Error(w, "player is not found", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(player.Buildings)
}

// 施設アップグレード
func UpgradeBuildingHandler(w http.ResponseWriter, r *http.Request) {
	var req UpgradeRequest
	//リクエストパラメータのチェック
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid Request", http.StatusBadRequest)
		return
	}

	//プレイヤー情報を取得
	player := storage.GetPlayer(req.PlayerID)
	if player == nil {
		http.Error(w, "player is not found", http.StatusBadRequest)
		return
	}

	//プレイヤーの建物情報を取得
	var building *models.Building
	for i := range player.Buildings {
		if player.Buildings[i].ID == req.BuildingID {
			building = &player.Buildings[i]
			break
		}
	}

	if building == nil {
		http.Error(w, "building not found", http.StatusBadRequest)
		return
	}

	//アップグレードのコスト計算
	cost := building.UpgradeCost()
	//コストの消費
	switch building.ResourceType {
	case "wood":
		if player.Resources.Wood < cost {
			http.Error(w, "resources not enough", http.StatusBadRequest)
			return
		}

		player.Resources.Wood -= cost

	case "stone":
		if player.Resources.Stone < cost {
			http.Error(w, "resources not enough", http.StatusBadRequest)
			return
		}
		player.Resources.Stone -= cost

	case "gold":
		if player.Resources.Gold < cost {
			http.Error(w, "resources not enough", http.StatusBadRequest)
			return
		}

		player.Resources.Gold -= cost
	}

	//施設のレベルアップ
	building.Level += 1
	storage.SavePlayer(player)

	json.NewEncoder(w).Encode(building)
}
