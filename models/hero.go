package models

type HeroBase struct {
	Name      string `json:"name"`
	Quality   int    `json:"quality"`
	Zhu       int    `json:"zhu"`
	Fu        int    `json:"fu"`
	SkillName string `json:"skillName" gorm:"column:skillName"`
	Img       string `json:"img"`
}

type HeroData struct {
	Effect string `json:"effect"`
}

type HeroAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	HeroBase
	Data []HeroData `json:"data"`
}

type HeroSelect struct {
	ID int `json:"id" gorm:"primaryKey"`
	HeroBase
	Data string `json:"data"`
}

type HeroAddParams struct {
	HeroBase
	Data []HeroData `json:"data"`
}

type HeroAddData struct {
	Data []HeroAddParams `json:"data"`
}

type HeroAddObj struct {
	HeroBase
	Data string `json:"data"`
}

type ShardBase struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Quality   int    `json:"quality"`
	LevelData string `json:"levelData" gorm:"column:levelData"`
	SkillData string `json:"skillData" gorm:"column:skillData"`
}

type ShardUpdateParams struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	SkillData string `json:"skillData" gorm:"column:skillData"`
}
