package models

type LogBase struct {
	Name string `json:"name"`
	Time string `json:"time"`
}

type LogAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	LogBase
}

func (LogAll) TableName() string { return "log" }

// 兼容旧命名
type LogList = LogAll
type LogAdd = LogBase
