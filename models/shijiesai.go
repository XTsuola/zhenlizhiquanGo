package models

type XuanshouInfo struct {
	Name string `json:"name"`
	Kedu int    `json:"kedu" gorm:"column:kedu"`
	Hero []int  `json:"hero" gorm:"column:hero"`
}

type ShijiesaiBase struct {
	No          int    `json:"no"`
	Info        string `json:"info"`
	ShengfuList string `json:"shengfuList" gorm:"column:shengfuList"`
}

type ShijiesaiAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	ShijiesaiBase
}

type ShijiesaiAdd struct {
	No          int `json:"no"`
	AInfo       XuanshouInfo
	BInfo       XuanshouInfo
	ShengfuList []int `json:"shengfuList" gorm:"column:shengfuList"`
}

type ShijiesaiList struct {
	ID int `json:"id" gorm:"primaryKey"`
	ShijiesaiAdd
}

type ShijiesaiAddData struct {
	Data []ShijiesaiList `json:"data"`
}
