package handlers

import (
	"net/http"
	"solitaire-serve-api/internal/models"
	"solitaire-serve-api/storage"
	"time"

	"github.com/gin-gonic/gin"
)

type BuildRequest struct {
	BuildingID int `json:"building_id"` //建物名
}

type UpgradeRequest struct {
	BuildingID int `json:"building_id"` //建物ID
}

// @Summary 施設の建築を実行する
// @Description プレイヤーが所持している資源を消費して施設の建築を実行する
// @Tags facility
// @Accept json
// @Produce json
// @Param body body BuildRequest true "PlayerID" "Name"
// @Success 200 {object} models.Building "facilities"
// @Failure 400 {string} string "invalid request or player is not found or resources not enough"
// @Router /facility/create [post]
func FacilityHandler(c *gin.Context) {
	var req BuildRequest
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

	//初期レベルの建設コスト計算
	tmpBuilding := models.Building{BuildingID: req.BuildingID, Level: 1}
	cost := tmpBuilding.UpgradeCost()
	//コストの消費と施設ごとの生産量を設定
	production := 1
	resource_type := 1
	switch tmpBuilding.ResourceType {
	case 1:
		if player.Resources.Wood < cost {
			c.JSON(http.StatusBadRequest, gin.H{"error": "resources not enough"})
			return
		}
		production = 5
		resource_type = 1

		player.Resources.Wood -= cost

	case 2:
		if player.Resources.Stone < cost {
			c.JSON(http.StatusBadRequest, gin.H{"error": "resources not enough"})
			return
		}
		player.Resources.Stone -= cost

		production = 3
		resource_type = 2

	case 3:
		if player.Resources.Gold < cost {
			c.JSON(http.StatusBadRequest, gin.H{"error": "resources not enough"})
			return
		}

		player.Resources.Gold -= cost

		production = 2
		resource_type = 3
	}

	//建築する建物情報をまとめる
	facilities := models.Building{
		BuildingID:    req.BuildingID,
		Level:         1,
		Position:      len(player.Buildings),
		Production:    production,
		ResourceType:  resource_type,
		LastCollected: time.Now(),
	}

	//リストに追加した上で保存実施
	player.Buildings = append(player.Buildings, facilities)
	storage.SavePlayer(player)

	c.JSON(http.StatusOK, gin.H{"facilities": facilities})
}

// @Summary 施設一覧
// @Description プレイヤーが建築済みの施設を一覧で取得する
// @Tags facility
// @Accept json
// @Produce json
// @Success 200 {object} models.Building "facilities"
// @Failure 400 {string} string "invalid request or player is not found"
// @Router /facility/list [post]
func FacilityListHandler(c *gin.Context) {
	player := storage.GetPlayer(playerId)
	if player == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "player is not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"facilities": player.Buildings})
}

// @Summary 施設アップグレード
// @Description 施設のレベルアップを行う
// @Tags facility
// @Accept json
// @Produce json
// @Param body body FacilityListRequest true "PlayerID" "BuildingID"
// @Success 200 {object} models.Building "facilities"
// @Failure 400 {string} string "invalid request or player is not found"
// @Router /facility/upgrade [post]
func UpgradeFacilityHandler(c *gin.Context) {
	var req UpgradeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//必要な情報がないのでエラー
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	//プレイヤー情報を取得
	player := storage.GetPlayer(playerId)
	if player == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "player is not found"})
		return
	}

	//プレイヤーの建物情報を取得
	var building *models.Building
	for i := range player.Buildings {
		if player.Buildings[i].BuildingID == req.BuildingID {
			building = &player.Buildings[i]
			break
		}
	}

	if building == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "building not found"})
		return
	}

	//アップグレードのコスト計算
	cost := building.UpgradeCost()
	//コストの消費
	switch building.ResourceType {
	case 1:
		if player.Resources.Wood < cost {
			c.JSON(http.StatusBadRequest, gin.H{"error": "resources not enough"})
			return
		}

		player.Resources.Wood -= cost

	case 2:
		if player.Resources.Stone < cost {
			c.JSON(http.StatusBadRequest, gin.H{"error": "resources not enough"})
			return
		}
		player.Resources.Stone -= cost

	case 3:
		if player.Resources.Gold < cost {
			c.JSON(http.StatusBadRequest, gin.H{"error": "resources not enough"})
			return
		}

		player.Resources.Gold -= cost
	}

	//施設のレベルアップ
	building.Level += 1
	storage.SavePlayer(player)

	c.JSON(http.StatusOK, gin.H{"facilities": building})
}
