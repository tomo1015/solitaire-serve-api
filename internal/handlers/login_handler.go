package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"solitaire-serve-api/internal/db"
	"solitaire-serve-api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ゲームトークンを生成して取得
func GetGameTokenHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PlatFormID int `json:"platformId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		//必要な情報がないのでエラー
		http.Error(w, "bad Request", http.StatusBadRequest)
		return
	}

	//プラットフォームIDからセッション情報を取得 または 生成
	session := GetSessionForPlatformId(req.PlatFormID)

	//ゲームトークンの生成
	gameToken := uuid.NewString()

	session.GameToken = gameToken
	if err := db.DB.Save(&session).Error; err != nil {
		http.Error(w, "update session is failed", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"game_token": session.GameToken})
}

// プラットフォームIDをもとに
// ユーザーIDを検索 または セッションテーブルの生成
func GetSessionForPlatformId(platformId int) models.Session {
	var session models.Session
	//プラットフォームIDでのセッション情報が存在する場合は
	//そのデータを返す
	result := db.DB.Where("platform_id = ?", platformId).First(&session)

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
