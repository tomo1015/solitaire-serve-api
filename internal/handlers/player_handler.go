package handlers

import (
	"encoding/json"
	"net/http"
	"solitaire-serve-api/internal/models"
	"solitaire-serve-api/storage"
)

func HandlePlayer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		player := storage.GetPlayer(id)
		if player == nil {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(player)
	case "POST":
		var p models.Player
		json.NewDecoder(r.Body).Decode(&p)
		storage.SavePlayer(&p)
		w.WriteHeader(http.StatusCreated)
	}
}
