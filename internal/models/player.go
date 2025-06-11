package models

type Resources struct {
	Wood  int `json:"wood"`  //木材
	Stone int `json:"stone"` //石材
	Gold  int `json:"gold"`  //金
}

type Player struct {
	ID        string     //ユーザーID
	Name      string     //ユーザー名
	Resources Resources  `json:"resources"` //資源情報
	Soldiers  int        //兵士情報
	Village   string     //村情報
	Buildings []Building `json:"building"` //建物リスト
}
