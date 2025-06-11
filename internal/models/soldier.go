package models

type Soldier struct {
	ID          string `json:"id"`           //兵士ID
	Type        string `json:"type"`         //兵士タイプ
	Level       int    `json:"level"`        //兵士レベル
	Quantity    int    `json:"quantity"`     //所持数
	Training    bool   `json:"training"`     //訓練中かどうか
	TrainingEnd int64  `json:"training_end"` //訓練終了のタイムスタンプ(unix秒)
}
