package storage

import (
	"encoding/json"
	"log"
	"os"
	"solitaire-serve-api/internal/db"
	"solitaire-serve-api/internal/models"
)

type rawDefensePoint struct {
	Name       string           `json:"name"`
	Location   models.WorldMap  `json:"location"`
	Type       string           `json:"type"`
	Difficulty int              `json:"difficulty"`
	Soldiers   []models.Soldier `json:"soldiers"`
	Loot       map[string]int   `json:"loot"`
}

func LoadDefensePointFromJson(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("failed load json file", err)
		return err
	}

	var raws []rawDefensePoint
	if err := json.Unmarshal(data, &raws); err != nil {
		log.Printf("failed rawDefensePoint", err)
		return err
	}

	for _, r := range raws {
		lootBytes, err := json.Marshal(r.Loot)
		if err != nil {
			log.Printf("loot material error", err)
			continue
		}

		soldiers := make([]*models.Soldier, len(r.Soldiers))
		for i := range r.Soldiers {
			soldiers[i] = &r.Soldiers[i]
		}

		dp := models.DefensePoint{
			Name:         r.Name,
			LocationType: r.Type,
			Difficulty:   r.Difficulty,
			Location:     r.Location,
			LootJson:     string(lootBytes),
			Soldiers:     soldiers,
		}

		if err := db.DB.Create(&dp).Error; err != nil {
			log.Printf("Failed to save defense point %s: %v", r.Name, err)
			continue
		}

		// 兵士情報も登録
		for _, s := range r.Soldiers {
			soldier := models.Soldier{
				ID:       dp.ID,
				Type:     s.Type,
				Level:    s.Level,
				Quantity: s.Quantity,
			}
			if err := db.DB.Create(&soldier).Error; err != nil {
				log.Printf("Failed to save soldier for %s: %v", r.Name, err)
			}
		}
	}

	return nil
}
