package models

type LogList struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Time string `json:"time"`
}

type LogAdd struct {
	Name string `json:"name"`
	Time string `json:"time"`
}
