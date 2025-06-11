package storage

import (
	"solitaire-serve-api/internal/models"
	"sync"
)

var Players = make(map[string]*models.Player)
var mu sync.Mutex

func GetPlayer(id string) *models.Player {
	mu.Lock()
	defer mu.Unlock()
	return Players[id]
}

func SavePlayer(p *models.Player) {
	mu.Lock()
	defer mu.Unlock()
	Players[p.ID] = p
}
