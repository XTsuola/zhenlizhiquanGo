package models

type CardBase struct {
	Name    string `json:"name"`
	Zhenyin int    `json:"zhenyin"`
	Quality int    `json:"quality"`
	Cost    int    `json:"cost"`
	Type    int    `json:"type"`
	Img     string `json:"img"`
	Grade   string `json:"grade"`
	Tag     string `json:"tag"`
}

type CardData struct {
	Attack int    `json:"attack"`
	Life   int    `json:"life"`
	Effect string `json:"effect"`
}

type CardAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	CardBase
	Data []CardData `json:"data"`
}

type CardSelect struct {
	ID int `json:"id" gorm:"primaryKey"`
	CardBase
	Data string `json:"data"`
}

type CardAddParams struct {
	CardBase
	Data []CardData `json:"data"`
}

type CardAddData struct {
	Data []CardAddParams `json:"data"`
}

type CardAddObj struct {
	CardBase
	Data string `json:"data"`
}

type CardUpdateGradeParams struct {
	ID    int   `json:"id" gorm:"primaryKey"`
	Grade []int `json:"grade"`
}

type CardUpdateGradeListParams struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Grade string `json:"grade"`
}

type CardUpdateGradeListData struct {
	Data []CardUpdateGradeListParams `json:"data"`
}

type CardUpdateTagParams struct {
	ID  int   `json:"id" gorm:"primaryKey"`
	Tag []int `json:"tag"`
}

type CardUpdateTag struct {
	ID  int    `json:"id" gorm:"primaryKey"`
	Tag string `json:"tag"`
}

type CardUpdateTagListParams struct {
	ID  int    `json:"id" gorm:"primaryKey"`
	Tag string `json:"tag"`
}

type CardUpdateTagListData struct {
	Data []CardUpdateTagListParams `json:"data"`
}
