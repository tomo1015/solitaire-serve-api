package storage

import (
	"encoding/json"
	"log"
	"os"
	"solitaire-serve-api/internal/db"
	"solitaire-serve-api/internal/models"
)

type rawDefensePoint struct {
	ID           int                    `json:"point_id"`
	Name         string                 `json:"name"`
	LocationX    int                    `json:"locationX"`
	LocationY    int                    `json:"locationY"`
	NPCName      string                 `json:"npc_name"`
	Soldiers     []models.BattleSoldier `json:"soldiers"`
	LocationType string                 `json:"type"`
	Loot         []models.Resources     `json:"loot"`
	Difficulty   int                    `json:"difficulty"`
}

type rawFacility struct {
	facilityID   int    `json:"facility_id"`
	Name         string `json:"name"`
	Level        int    `json:"level"`
	Production   int    `json:"production"`
	ResourceType int    `json:"resource_type"`
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
		soldiers := make([]*models.BattleSoldier, len(r.Soldiers))
		for i := range r.Soldiers {
			soldiers[i] = &r.Soldiers[i]
		}

		resources := make([]*models.Resources, len(r.Loot))
		for i := range r.Loot {
			resources[i] = &r.Loot[i]
		}

		dp := models.DefensePoint{
			ID:           r.ID,
			Name:         r.Name,
			LocationX:    r.LocationX,
			LocationY:    r.LocationY,
			NPCName:      r.NPCName,
			Soldiers:     soldiers,
			Loot:         resources,
			LocationType: r.LocationType,
			Difficulty:   r.Difficulty,
		}

		if err := db.DB.Create(&dp).Error; err != nil {
			log.Printf("Failed to save defense point %s: %v", r.Name, err)
			continue
		}

		// 兵士情報も登録
		for _, s := range r.Soldiers {
			soldier := models.BattleSoldier{
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

func LoadFacilityFromJson(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("failed load json file", err)
		return err
	}

	var raws []rawFacility
	if err := json.Unmarshal(data, &raws); err != nil {
		log.Printf("failed rawDefensePoint", err)
		return err
	}

	for _, r := range raws {

		dp := models.Building{
			BuildingID:   r.facilityID,
			Level:        r.Level,
			Production:   r.Production,
			ResourceType: r.ResourceType,
		}

		if err := db.DB.Create(&dp).Error; err != nil {
			log.Printf("Failed to save facility %s: %v", r.Name, err)
			continue
		}
	}

	return nil
}
