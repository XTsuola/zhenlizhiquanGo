package models

type LogBase struct {
	Name string `json:"name"`
	Time string `json:"time"`
}

type LogAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	LogBase
}

// TableName 表名 log
func (LogAll) TableName() string { return "log" }

// LogList、LogAdd 兼容旧命名，分别等同 LogAll、LogBase
type LogList = LogAll
type LogAdd = LogBase
