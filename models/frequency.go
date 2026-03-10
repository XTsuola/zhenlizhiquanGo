package models

type FrequencyBase struct {
	Name     string `json:"name"`
	Qu       int    `json:"qu"`
	HeroId   int    `json:"heroId" gorm:"column:heroId"`
	HeroLife int    `json:"heroLife" gorm:"column:heroLife"`
	Cards    string `json:"cards"`
	Time     string `json:"time"`
}

type FrequencyAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	FrequencyBase
}

type FrequencyUpdate struct {
	ID int `json:"id" gorm:"primaryKey"`
	FrequencyBase
	Password string `json:"password"`
}

type FrequencyAddAll struct {
	Data []FrequencyBase `json:"data"`
}

type FrequencyPaddwordAll struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Password string `json:"password"`
}

type FrequencyPaddwordAdd struct {
	Password string `json:"password"`
}
