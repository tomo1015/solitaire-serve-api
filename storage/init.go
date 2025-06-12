package storage

import (
	"encoding/json"
	"log"
	"os"
	"solitaire-serve-api/internal/models"
)

func LoadDefensePointFromJson(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("failed to open defense point JSON: %v", err)
	}
	defer file.Close()

	var points []*models.DefensePoint
	if err := json.NewDecoder(file).Decode(&points); err != nil {
		log.Fatalf("failed to decode defense point JSON: %v", err)
	}

	DefensePoints = points
	log.Printf("Loaded %d defense points from %s", len(points), filepath)
}
