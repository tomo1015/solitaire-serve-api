package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"solitaire-serve-api/internal/models"
	"solitaire-serve-api/storage"
	"time"

	"github.com/gin-gonic/gin"
)

type TrainRequest struct {
	Type     string `json:"type"`     //兵士タイプ
	Quantity int    `json:"quantity"` //訓練数
}

// @Summary 兵士一覧
// @Description 訓練できる兵士の一覧
// @Tags soldier
// @Accept json
// @Produce json
// @Param body body TrainListRequest true "PlayerID"
// @Success 200 {object}
// @Failure 400 {string} string "invalid request or player is not found"
// @Router /soldier/list [post]
func GetSoldiersHandler(c *gin.Context) {
	//ストレージからプレイヤー情報を取得
	player := storage.GetPlayer(playerId)
	if player == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "player is not found"})
		return
	}
}

// @Summary 兵士一覧
// @Description 訓練できる兵士の一覧
// @Tags soldier
// @Accept json
// @Produce json
// @Param body body TrainListRequest true "PlayerID"
// @Success 200 {object} true "is_ok"
// @Failure 400 {string} string "invalid request or player is not found"
// @Router /soldier/list [post]
func TrainHandler(c *gin.Context) {
	var req TrainRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//必要な情報がないのでエラー
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	//ストレージからプレイヤー情報を取得
	player := storage.GetPlayer(playerId)
	if player == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "player is not found"})
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

	c.JSON(http.StatusOK, gin.H{"is_ok": true})
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
