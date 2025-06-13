package models

// 防衛拠点の構造体
type DefensePoint struct {
	ID           int              `json:"id"`
	Name         string           `json:"name"`       //拠点名
	Location     WorldMap         `json:"location"`   //拠点の座標
	NPCName      string           `json:"npc_name"`   //NPC名
	Soldiers     []*BattleSoldier `json:"soldiers"`   //防衛兵士
	Loot         Resources        `json:"loot"`       //獲得できる資源
	LootJson     string           `json:"-"`          //Jsonからのデータを保存する
	LocationType string           `json:"type"`       //資源産出地域タイプ
	Difficulty   int              `json:"difficulty"` //拠点難易度
}
