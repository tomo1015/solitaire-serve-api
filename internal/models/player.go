package models

type Resources struct {
	ID    int `json:"id"`
	Wood  int `json:"wood"`  //木材
	Stone int `json:"stone"` //石材
	Gold  int `json:"gold"`  //金
}

type Player struct {
	ID        string     //ユーザーID
	Name      string     `json:"name"`      //ユーザー名
	Resources Resources  `json:"resources"` //資源情報
	Soldiers  []*Soldier `json:"soldiers"`  //兵士情報
	Point     int        `json:"point"`     //戦果ポイント
	Village   string     //村情報
	Buildings []Building `json:"building"` //建物リスト
}

// 倉庫レベル用取得関数
func (player *Player) GetWarehouseLevel() int {
	for _, b := range player.Buildings {
		if b.Name == "倉庫" {
			return b.Level
		}
	}
	return 0
}

// 兵士情報を取得する関数
func (player *Player) FindSoldier(soldierID int) *Soldier {
	for i := range player.Soldiers {
		if player.Soldiers[i].ID == soldierID {
			return player.Soldiers[i] //この時点でポインタなのでそのまま書くことができる
		}
	}
	return nil
}
