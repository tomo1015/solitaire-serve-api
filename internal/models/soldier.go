package models

type Soldier struct {
	ID       int    `json:"id"`       //兵士ID
	Type     string `json:"type"`     //兵士タイプ
	Level    int    `json:"level"`    //兵士レベル
	Quantity int    `json:"quantity"` //所持数
	//訓練用パラメータ
	Training    bool  `json:"training"`     //訓練中かどうか
	TrainingEnd int64 `json:"training_end"` //訓練終了のタイムスタンプ(unix秒)
}

// 戦闘用パラメータ
type BattleSoldier struct {
	ID       int     `json:"id"`        //兵士ID
	Type     string  `json:"type"`      //兵士タイプ
	Level    int     `json:"level"`     //兵士レベル
	Quantity int     `json:"quantity"`  //バトルに使う兵士数
	Attack   int     `json:"attack"`    //攻撃力
	Defense  int     `json:"defense"`   //防御力
	HitRate  float64 `json:"hit_rate"`  //命中率(0～1)
	CritRate float64 `json:"crit_rate"` //クリティカル確率(0～1)
}
