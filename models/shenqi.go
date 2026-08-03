package models

import "go_project/utils"

type ShenqiBase struct {
	Name    string `json:"name"`
	Zhenyin int    `json:"zhenyin"`
	Quality int    `json:"quality"`
	Type    int    `json:"type"`
	Img     string `json:"img"`
	Bonus   string `json:"bonus"`
}

// ShenqiData 神器效果（存在 Data JSON 中）
type ShenqiData struct {
	Effect string `json:"effect"`
}

// ShenqiSelect DB 读取形态
type ShenqiSelect struct {
	ID int `json:"id" gorm:"primaryKey"`
	ShenqiBase
	Data string `json:"data"`
}

func (ShenqiSelect) TableName() string { return "shenqi" }

// ShenqiAll API 输出形态
type ShenqiAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	ShenqiBase
	Data []ShenqiData `json:"data"`
}

func (item ShenqiSelect) ToAll() ShenqiAll {
	return ShenqiAll{
		ID:         item.ID,
		ShenqiBase: item.ShenqiBase,
		Data:       utils.StringToArr[ShenqiData](item.Data),
	}
}

type ShenqiAddParams struct {
	ShenqiBase
	Data []ShenqiData `json:"data"`
}

func (p ShenqiAddParams) ToObj() ShenqiAddObj {
	return ShenqiAddObj{
		ShenqiBase: p.ShenqiBase,
		Data:       utils.ArrToString(p.Data),
	}
}

type ShenqiAddData struct {
	Data []ShenqiAddParams `json:"data"`
}

type ShenqiAddObj struct {
	ShenqiBase
	Data string `json:"data"`
}

func (ShenqiAddObj) TableName() string { return "shenqi" }
