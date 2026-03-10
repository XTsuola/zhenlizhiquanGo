package models

type ShenqiBase struct {
	Name    string `json:"name"`
	Zhenyin int    `json:"zhenyin"`
	Quality int    `json:"quality"`
	Type    int    `json:"type"`
	Img     string `json:"img"`
}

type ShenqiData struct {
	Effect string `json:"effect"`
}

type ShenqiAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	ShenqiBase
	Data []ShenqiData `json:"data"`
}

type ShenqiSelect struct {
	ID int `json:"id" gorm:"primaryKey"`
	ShenqiBase
	Data string `json:"data"`
}

type ShenqiAddParams struct {
	ShenqiBase
	Data []ShenqiData `json:"data"`
}

type ShenqiAddData struct {
	Data []ShenqiAddParams `json:"data"`
}

type ShenqiAddObj struct {
	ShenqiBase
	Data string `json:"data"`
}
