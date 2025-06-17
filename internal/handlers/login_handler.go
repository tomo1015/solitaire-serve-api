package handlers

import (
	"errors"
	"net/http"
	"solitaire-serve-api/internal/db"
	"solitaire-serve-api/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GetGameTokenRequest struct {
	PlatFormID int `json:"platformId"`
}

var playerId int //ユーザーID

// @Summary ゲームトークン取得
// @Description プラットフォームIDに対応した接続情報を作成し、ゲームトークンを生成する
// @Tags login
// @Accept json
// @Produce json
// @Param body body GetGameTokenRequest true "Platform ID"
// @Success 200 {object} map[string]string "game_token"
// @Failure 400 {string} string "bad Request or update Failed"
// @Router /getGameToken [post]
func GetGameTokenHandler(c *gin.Context) {
	var req GetGameTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//必要な情報がないのでエラー
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	//プラットフォームIDからセッション情報を取得 または 生成
	session := GetSessionForPlatformId(req.PlatFormID)

	//ゲームトークンの生成
	gameToken := uuid.NewString()

	session.GameToken = gameToken
	if err := db.DB.Save(&session).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update failed"})
		return
	}

	playerId = session.UserID
	c.JSON(http.StatusOK, gin.H{"game_token": session.GameToken})
}

// プラットフォームIDをもとに
// ユーザーIDを検索 または セッションテーブルの生成
func GetSessionForPlatformId(platformId int) models.Session {
	var session models.Session
	//プラットフォームIDでのセッション情報が存在する場合は
	//そのデータを返す
	result := db.DB.Where("PlatFormID = ?", platformId).First(&session)

	if result.Error == nil {
		//データがあるのでそのSession情報を返す
		return session
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Session{}
	}

	//データがないので新規生成する
	newSession := models.Session{
		PlatFormID: platformId,
	}

	//とりあえずデータを生成する
	if err := db.DB.Create(&newSession).Error; err != nil {
		return models.Session{}
	}

	//生成されたuser_idを取得
	playerId = newSession.UserID

	//シャードキー生成
	newSession.ShardKey = GetShardKeyForUserId(newSession.UserID)
	if err := db.DB.Save(&newSession).Error; err != nil {
		return models.Session{}
	}

	return newSession
}

// シャードキーの取得
func GetShardKeyForUserId(userId int) int {
	const numShards = 4 //4シャード分準備
	return int(userId % numShards)
}
