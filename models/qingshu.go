package models

type QingshuUserData struct {
	ID        int    `json:"id" `
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

type QingshuMapParams struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	CardPile string `json:"cardPile" gorm:"column:cardPile"`
	DisPile  string `json:"disPile" gorm:"column:disPile"`
	UserData string `json:"userData" gorm:"column:userData"`
	QingshuMapBase
}

type QingshuMapData struct {
	ID       int               `json:"id" gorm:"primaryKey"`
	CardPile []int             `json:"cardPile" gorm:"column:cardPile"`
	DisPile  []int             `json:"disPile" gorm:"column:disPile"`
	UserData []QingshuUserData `json:"userData" gorm:"column:userData"`
	QingshuMapBase
}

type QingshuDisCard struct {
	MyId    int `json:"myId" gorm:"column:myId"`
	Pai     int `json:"pai"`
	YourPai int `json:"yourPai" gorm:"column:yourPai"`
	Index   int `json:"index"`
}

type Message struct {
	Type    int `json:"type" gorm:"column:type"`
	UserId  int `json:"userId" gorm:"column:userId"`
	Pai     int `json:"pai"`
	YourPai int `json:"yourPai" gorm:"column:yourPai"`
	Index   int `json:"index"`
}

type ReturnMessage struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UsernameUpdate struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
