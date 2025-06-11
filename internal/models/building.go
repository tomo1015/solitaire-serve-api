package models

type Building struct {
	ID       string //一意なID（建設時に発行）
	Name     string //施設名
	Level    int    //施設レベル（初期1)
	Position int    //建物の位置
}
