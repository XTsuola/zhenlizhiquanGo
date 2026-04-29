package models

type MemberBase struct {
	Name   string  `json:"name"`
	Score  float64 `json:"score"`
	Title  string  `json:"title"`
	Remark string  `json:"remark"`
}

type MemberAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	MemberBase
}

type MemberAddData struct {
	Data []MemberBase `json:"data"`
}
