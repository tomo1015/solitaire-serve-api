package models

type Player struct {
	ID        string     //ユーザーID
	Name      string     //ユーザー名
	Resources int        //資源情報
	Soldiers  int        //兵士情報
	Village   string     //村情報
	Buildings []Building //建物リスト
}
