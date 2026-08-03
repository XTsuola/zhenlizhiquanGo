package models

type FrequencyBase struct {
	Name     string `json:"name"`
	Qu       int    `json:"qu"`
	HeroId   int    `json:"heroId" gorm:"column:heroId"`
	HeroLife int    `json:"heroLife" gorm:"column:heroLife"`
	Cards    string `json:"cards"` // JSON 字符串
	Time     string `json:"time"`
}

type FrequencyAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	FrequencyBase
}

func (FrequencyAll) TableName() string { return "frequency" }

type FrequencyUpdate struct {
	ID int `json:"id"`
	FrequencyBase
	Password string `json:"password"` // 非表字段
}

type FrequencyAddAll struct {
	Data []FrequencyBase `json:"data"`
}
