package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"solitaire-serve-api/internal/models"
	"solitaire-serve-api/storage"
	"time"
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
		Type:        req.Type,
		Level:       1,
		Quantity:    req.Quantity,
		Training:    true,
		TrainingEnd: now + trainDuration,
	}

	//DBに追加
	db, err := sql.Open("sqlite3", "db/game.sqlite") //DBは指定フォルダにあるもの
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := insertSoldier(db, soldier); err != nil {
		panic(err)
	}

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

// DBに情報を追加
func insertSoldier(db *sql.DB, soldier *models.Soldier) error {
	result, err := db.Exec(
		`INSERT INTO soldiers (type, level, quantity, training, training_end) VALUES (?, ?, ?, ?, ?)`,
		soldier.Type, soldier.Level, soldier.Quantity, soldier.Training, soldier.TrainingEnd,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	soldier.ID = int(id) // 自動採番されたIDを構造体にセット
	fmt.Printf("Inserted soldier with ID: %d\n", soldier.ID)
	return nil
}
