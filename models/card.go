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

func (CardSelect) TableName() string { return "card" }

// CardAll API 输出形态（Data 已反序列化）
type CardAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	CardBase
	Data []CardData `json:"data"`
}

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

func (p CardAddParams) ToObj() CardAddObj {
	return CardAddObj{
		CardBase: p.CardBase,
		Data:     utils.ArrToString(p.Data),
	}
}

// CardAddData 批量新增请求
type CardAddData struct {
	Data []CardAddParams `json:"data"`
}

// CardAddObj DB 写入形态
type CardAddObj struct {
	CardBase
	Data string `json:"data"`
}

func (CardAddObj) TableName() string { return "card" }

// CardUpdateGradeParams 单条更新 grade（请求为 []int）
type CardUpdateGradeParams struct {
	ID    int   `json:"id"`
	Grade []int `json:"grade"`
}

// CardUpdateGradeListParams 批量更新 grade（请求已是字符串）
type CardUpdateGradeListParams struct {
	ID    int    `json:"id"`
	Grade string `json:"grade"`
}

type CardUpdateGradeListData struct {
	Data []CardUpdateGradeListParams `json:"data"`
}

// CardUpdateTagParams 单条更新 tag（请求为 []int）
type CardUpdateTagParams struct {
	ID  int   `json:"id"`
	Tag []int `json:"tag"`
}

// CardUpdateTagListParams 批量更新 tag
type CardUpdateTagListParams struct {
	ID  int    `json:"id"`
	Tag string `json:"tag"`
}

type CardUpdateTagListData struct {
	Data []CardUpdateTagListParams `json:"data"`
}
