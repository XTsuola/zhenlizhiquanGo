package models

type SkinBase struct {
	CardId   int    `json:"cardId" gorm:"column:cardId"`
	Name     string `json:"name"`
	Zhenyin  int    `json:"zhenyin"`
	Cost     int    `json:"cost"`
	Skill    string `json:"skill"`
	Img      string `json:"img"`
	Shuxing  string `json:"shuxing"`
	Origin   string `json:"origin"`
	Remark   string `json:"remark"`
	Together string `json:"together"`
}

type SkinAddParams struct {
	SkinBase
	Effect []string `json:"effect"`
}

type SkinAddData struct {
	Data []SkinAddParams `json:"data"`
}

type SkinAddObj struct {
	SkinBase
	Effect string `json:"effect"`
}

type SkinSelect struct {
	ID int `json:"id" gorm:"primaryKey"`
	SkinBase
	Effect string `json:"effect"`
}

type SkinAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	SkinBase
	Effect []string `json:"effect"`
}

type SkinTogetherParams struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Together string `json:"together"`
}

type SkinUpdateData struct {
	Data []SkinTogetherParams `json:"data"`
}
