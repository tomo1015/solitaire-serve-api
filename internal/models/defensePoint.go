package models

// 防衛拠点の構造体
type DefensePoint struct {
	ID           int              `gorm:"primaryKey;autoIncrement",json:"point_id"`
	Name         string           `json:"name"` //拠点名
	LocationX    int              `json:"locationX"`
	LocationY    int              `json:"LocationY"`
	Soldiers     []*BattleSoldier `gorm:"foreignKey:ID"`
	NPCName      string           `json:"npc_name"`      //NPC名
	Loot         []*Resources     `gorm:"foreignKey:ID"` //獲得できる資源
	LocationType string           `json:"type"`          //資源産出地域タイプ
	Difficulty   int              `json:"difficulty"`    //拠点難易度
}
