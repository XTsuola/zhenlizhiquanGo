package models

import "go_project/utils"

type SkinBase struct {
	CardId  int    `json:"cardId" gorm:"column:cardId"`
	Name    string `json:"name"`
	Zhenyin int    `json:"zhenyin"`
	Cost    int    `json:"cost"`
	Skill   string `json:"skill"`
	Img     string `json:"img"`
	Shuxing string `json:"shuxing"`
	Origin  string `json:"origin"`
	Remark  string `json:"remark"`
}

// SkinSelect DB 读取形态（Effect 为 JSON 字符串）
type SkinSelect struct {
	ID int `json:"id" gorm:"primaryKey"`
	SkinBase
	Effect string `json:"effect"`
}

// TableName 表名 skin
func (SkinSelect) TableName() string { return "skin" }

// SkinAll API 输出形态
type SkinAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	SkinBase
	Effect []string `json:"effect"`
}

// ToAll 转为接口输出，把 Effect 从 JSON 字符串还原为字符串列表
func (item SkinSelect) ToAll() SkinAll {
	return SkinAll{
		ID:       item.ID,
		SkinBase: item.SkinBase,
		Effect:   utils.StringToArr[string](item.Effect),
	}
}

type SkinAddParams struct {
	SkinBase
	Effect []string `json:"effect"`
}

// ToObj 转为入库形态，把 Effect 序列化为 JSON 字符串
func (p SkinAddParams) ToObj() SkinAddObj {
	return SkinAddObj{
		SkinBase: p.SkinBase,
		Effect:   utils.ArrToString(p.Effect),
	}
}

type SkinAddObj struct {
	SkinBase
	Effect string `json:"effect"`
}

// TableName 表名 skin
func (SkinAddObj) TableName() string { return "skin" }
