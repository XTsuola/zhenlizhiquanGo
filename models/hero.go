package models

import "go_project/utils"

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

// HeroSelect DB 读取形态（Agent/Data 为 JSON 字符串）
type HeroSelect struct {
	ID int `json:"id" gorm:"primaryKey"`
	HeroBase
	Agent string `json:"agent"`
	Data  string `json:"data"`
}

func (HeroSelect) TableName() string { return "hero" }

// HeroAll API 输出形态
type HeroAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	HeroBase
	Agent []string   `json:"agent"`
	Data  []HeroData `json:"data"`
}

func (item HeroSelect) ToAll() HeroAll {
	return HeroAll{
		ID:       item.ID,
		HeroBase: item.HeroBase,
		Agent:    utils.StringToArr[string](item.Agent),
		Data:     utils.StringToArr[HeroData](item.Data),
	}
}

// HeroAddParams 单条新增请求（前端 agent 为 string[]）
type HeroAddParams struct {
	HeroBase
	Agent []string   `json:"agent"`
	Data  []HeroData `json:"data"`
}

func (p HeroAddParams) ToObj() HeroAddObj {
	return HeroAddObj{
		HeroBase: p.HeroBase,
		Agent:    utils.ArrToString(p.Agent),
		Data:     utils.ArrToString(p.Data),
	}
}

type HeroAddData struct {
	Data []HeroAddParams `json:"data"`
}

// HeroAddObj DB 写入形态
type HeroAddObj struct {
	HeroBase
	Agent string `json:"agent"`
	Data  string `json:"data"`
}

func (HeroAddObj) TableName() string { return "hero" }

// ---------- 碎片 ----------

type ShardBase struct {
	Quality   int    `json:"quality"`
	LevelData string `json:"levelData" gorm:"column:levelData"`
	SkillData string `json:"skillData" gorm:"column:skillData"`
}

type ShardAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	ShardBase
}

func (ShardAll) TableName() string { return "shard" }

type ShardUpdateParams struct {
	ID        int    `json:"id"`
	SkillData string `json:"skillData"`
}

type ShardAddData struct {
	Data []ShardBase `json:"data"`
}

type AgentUpdate struct {
	ID    int      `json:"id"`
	Agent []string `json:"agent"`
}
