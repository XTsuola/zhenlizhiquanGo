package models

import "go_project/utils"

// CardBase 卡牌基础字段（DB 列）
type CardBase struct {
	Name    string `json:"name"`
	Zhenyin int    `json:"zhenyin"`
	Quality int    `json:"quality"`
	Cost    int    `json:"cost"`
	Type    int    `json:"type"`
	Img     string `json:"img"`
	Grade   string `json:"grade"` // JSON 数组字符串，如 "[1,2]"
	Tag     string `json:"tag"`   // JSON 数组字符串
}

// CardData 卡牌等级数据（存在 Data JSON 中）
type CardData struct {
	Attack int    `json:"attack"`
	Life   int    `json:"life"`
	Effect string `json:"effect"`
}

// CardSelect DB 读取形态（Data 为 JSON 字符串）
type CardSelect struct {
	ID int `json:"id" gorm:"primaryKey"`
	CardBase
	Data string `json:"data"`
}

// TableName 表名 card
func (CardSelect) TableName() string { return "card" }

// CardAll API 输出形态（Data 已反序列化）
type CardAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	CardBase
	Data []CardData `json:"data"`
}

// ToAll 转为接口输出，把 Data 从 JSON 字符串还原为等级列表
func (item CardSelect) ToAll() CardAll {
	return CardAll{
		ID:       item.ID,
		CardBase: item.CardBase,
		Data:     utils.StringToArr[CardData](item.Data),
	}
}

// CardAddParams 单条新增请求
type CardAddParams struct {
	CardBase
	Data []CardData `json:"data"`
}

// ToObj 转为入库形态，把 Data 序列化为 JSON 字符串
func (p CardAddParams) ToObj() CardAddObj {
	return CardAddObj{
		CardBase: p.CardBase,
		Data:     utils.ArrToString(p.Data),
	}
}

// CardAddObj DB 写入形态
type CardAddObj struct {
	CardBase
	Data string `json:"data"`
}

// TableName 表名 card
func (CardAddObj) TableName() string { return "card" }

// CardUpdateGradeParams 单条更新 grade（请求为 []int）
type CardUpdateGradeParams struct {
	ID    int   `json:"id"`
	Grade []int `json:"grade"`
}

// CardUpdateTagParams 单条更新 tag（请求为 []int）
type CardUpdateTagParams struct {
	ID  int   `json:"id"`
	Tag []int `json:"tag"`
}
