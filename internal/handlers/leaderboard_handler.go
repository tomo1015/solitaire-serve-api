package handlers

import (
	"encoding/json"
	"net/http"
	"solitaire-serve-api/internal/leaderboard"
)

func HandleLeaderboard(w http.ResponseWriter, r *http.Request) {
	lb := leaderboard.GetLeaderboard()
	json.NewEncoder(w).Encode(lb)
}
