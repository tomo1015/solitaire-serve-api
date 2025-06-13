package models

type Attack struct {
	ID            string           `json:"id"`
	AttackerID    string           `json:"attacker_id"` //攻撃ユーザーID
	Target        WorldMap         `json:"target"`      //ワールドマップの座標
	BattleSoldier []*BattleSoldier `json:"soldiers"`    //派遣兵士の情報
	ExecuteAt     int64            `json:"execute_at"`  //攻撃実行時間
	Processed     bool             `json:"processed"`   //攻撃実行中かどうか
	Result        string           `json:"result"`      //攻撃結果
}
