package models

type SkinDiyBase struct {
	CardId string `json:"cardId" gorm:"column:cardId"`
	Name   string `json:"name"`
	Skill  string `json:"skill"`
	Effect string `json:"effect"`
	Reason string `json:"reason"`
	Remark string `json:"remark"`
}

type SkinDiyAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	SkinDiyBase
}

type SkinDiyUpdate struct {
	ID int `json:"id" gorm:"primaryKey"`
	SkinDiyBase
	Password string `json:"password"`
}

type SkinDiyAddData struct {
	Data []SkinDiyBase `json:"data"`
}

type CardDiyBase struct {
	Name     string `json:"name"`
	Zhenyin  int    `json:"zhenyin"`
	Cost     int    `json:"cost"`
	Quality  int    `json:"quality"`
	CardType int    `json:"cardType" gorm:"column:cardType"`
	Att      int    `json:"att"`
	Life     int    `json:"life"`
	Effect   string `json:"effect"`
	Img      string `json:"img"`
	Info     string `json:"info"`
	Remark   string `json:"remark"`
}

type CardDiyAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	CardDiyBase
}

type CardDiyUpdate struct {
	ID int `json:"id" gorm:"primaryKey"`
	CardDiyBase
	Password string `json:"password"`
}

type CardDiyAddData struct {
	Data []CardDiyBase `json:"data"`
}
