package models

// SkinDiyBase 皮肤 DIY（cardId 存字符串，可表示多卡关联）
type SkinDiyBase struct {
	CardId string `json:"cardId" gorm:"column:cardId"`
	Name   string `json:"name"`
	Skill  string `json:"skill"`
	Effect string `json:"effect"`
	Reason string `json:"reason"`
	Remark string `json:"remark"`
}

type SkinDiyAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	SkinDiyBase
}

func (SkinDiyAll) TableName() string { return "skin_diy" }

type SkinDiyUpdate struct {
	ID int `json:"id"`
	SkinDiyBase
	Password string `json:"password"` // 管理员或临时密码，非表字段
}

type SkinDiyAddData struct {
	Data []SkinDiyBase `json:"data"`
}

// CardDiyBase 卡牌 DIY
type CardDiyBase struct {
	Name     string `json:"name"`
	Zhenyin  int    `json:"zhenyin"`
	Cost     int    `json:"cost"`
	Quality  int    `json:"quality"`
	CardType int    `json:"cardType" gorm:"column:cardType"`
	Att      int    `json:"att"`
	Life     int    `json:"life"`
	Effect   string `json:"effect"`
	Img      string `json:"img"`
	Info     string `json:"info"`
	Remark   string `json:"remark"`
}

type CardDiyAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	CardDiyBase
}

func (CardDiyAll) TableName() string { return "card_diy" }

type CardDiyUpdate struct {
	ID int `json:"id"`
	CardDiyBase
	Password string `json:"password"` // 非表字段
}

type CardDiyAddData struct {
	Data []CardDiyBase `json:"data"`
}
