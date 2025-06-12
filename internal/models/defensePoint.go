package models

//防衛拠点の構造体
type DefensePoint struct {
	Location     WorldMap   `json:"location"` //拠点の座標
	NPCName      string     `json:"npc_name"` //NPC名
	Soldiers     []*Soldier `json:"soldiers"` //防衛兵士
	Loot         Resources  `json:"loot"`     //獲得できる資源
	LocationType string     `json:"type"`     //資源産出地域タイプ
}
