package models

type SkinDiyBase struct {
	CardId int    `json:"cardId" gorm:"column:cardId"`
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
