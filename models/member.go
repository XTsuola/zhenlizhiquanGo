package models

type MemberBase struct {
	Name     string  `json:"name"`
	Donation float64 `json:"donation"`
	Score    float64 `json:"score"`
	Title    string  `json:"title"`
	Remark   string  `json:"remark"`
}

type MemberAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	MemberBase
}

func (MemberAll) TableName() string { return "member" }

type MemberAddData struct {
	Data []MemberBase `json:"data"`
}
