package handlers

import (
	"encoding/json"
	"net/http"
	"solitaire-serve-api/internal/models"
	"solitaire-serve-api/internal/util"
	"solitaire-serve-api/storage"
)

func HandlePlayer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		player := storage.GetPlayer(playerId)
		if player == nil {
			http.NotFound(w, r)
			return
		}

		//プレイヤー取得時に資源を自動加算する
		util.CollectResources(player)
		storage.SavePlayer(player)

		json.NewEncoder(w).Encode(player)
	case "POST":
		var p models.Player
		json.NewDecoder(r.Body).Decode(&p)
		storage.SavePlayer(&p)
		w.WriteHeader(http.StatusCreated)
	}
}
