package models

import "go_project/utils"

// QingshuUserData 对局中的玩家状态（嵌在 userData JSON）
type QingshuUserData struct {
	ID        int    `json:"id"`
	HandCards []int  `json:"handCards"`
	DisCards  []int  `json:"disCards"`
	UserName  string `json:"userName"`
	Status    int    `json:"status"`
}

type QingshuMapBase struct {
	Round  int    `json:"round"`
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// QingshuMapParams DB 存储形态（牌堆/用户为 JSON 字符串）
type QingshuMapParams struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	CardPile string `json:"cardPile" gorm:"column:cardPile"`
	DisPile  string `json:"disPile" gorm:"column:disPile"`
	UserData string `json:"userData" gorm:"column:userData"`
	QingshuMapBase
}

func (QingshuMapParams) TableName() string { return "qingshu" }

// QingshuMapData 运行时 / API 形态
type QingshuMapData struct {
	ID       int               `json:"id" gorm:"primaryKey"`
	CardPile []int             `json:"cardPile"`
	DisPile  []int             `json:"disPile"`
	UserData []QingshuUserData `json:"userData"`
	QingshuMapBase
}

func (p QingshuMapParams) ToData() QingshuMapData {
	return QingshuMapData{
		ID:       p.ID,
		CardPile: utils.StringToArr[int](p.CardPile),
		DisPile:  utils.StringToArr[int](p.DisPile),
		UserData: utils.StringToArr[QingshuUserData](p.UserData),
		QingshuMapBase: QingshuMapBase{
			Round:  p.Round,
			Status: p.Status,
			Msg:    p.Msg,
		},
	}
}

func (d QingshuMapData) ToUpdateMap() map[string]interface{} {
	return map[string]interface{}{
		"round":    d.Round,
		"status":   d.Status,
		"msg":      d.Msg,
		"cardPile": utils.ArrToString(d.CardPile),
		"disPile":  utils.ArrToString(d.DisPile),
		"userData": utils.ArrToString(d.UserData),
	}
}

// QingshuDisCard 出牌相关参数（预留）
type QingshuDisCard struct {
	MyId    int `json:"myId"`
	Pai     int `json:"pai"`
	YourPai int `json:"yourPai"`
	Index   int `json:"index"`
}

// Message WebSocket 入站消息
type Message struct {
	Type    int `json:"type"`
	UserId  int `json:"userId"`
	Pai     int `json:"pai"`
	YourPai int `json:"yourPai"`
	Index   int `json:"index"`
}

// ReturnMessage WebSocket 出站消息
type ReturnMessage struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// UsernameUpdate 修改昵称（Password 实际为玩家槽位 "1"/"2"）
type UsernameUpdate struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
