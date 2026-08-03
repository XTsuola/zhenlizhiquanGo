package models

import "go_project/utils"

// XuanshouInfo 选手信息（嵌在 shijiesai.info JSON 中，非独立表列）
type XuanshouInfo struct {
	Name string `json:"name"`
	Kedu int    `json:"kedu"`
	Hero []int  `json:"hero"`
}

// ShijiesaiBase DB 写入/存储形态
type ShijiesaiBase struct {
	No          int    `json:"no"`
	Info        string `json:"info"`                                  // [AInfo, BInfo] JSON
	ShengfuList string `json:"shengfuList" gorm:"column:shengfuList"` // []int JSON
}

func (ShijiesaiBase) TableName() string { return "shijiesai" }

// ShijiesaiAll DB 完整行
type ShijiesaiAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	ShijiesaiBase
}

func (ShijiesaiAll) TableName() string { return "shijiesai" }

// ShijiesaiAdd API 新增请求
type ShijiesaiAdd struct {
	No          int          `json:"no"`
	AInfo       XuanshouInfo `json:"AInfo"`
	BInfo       XuanshouInfo `json:"BInfo"`
	ShengfuList []int        `json:"shengfuList"`
}

func (p ShijiesaiAdd) ToBase() ShijiesaiBase {
	return ShijiesaiBase{
		No:          p.No,
		Info:        utils.ArrToString([]XuanshouInfo{p.AInfo, p.BInfo}),
		ShengfuList: utils.ArrToString(p.ShengfuList),
	}
}

// ShijiesaiList API 列表/更新形态
type ShijiesaiList struct {
	ID int `json:"id"`
	ShijiesaiAdd
}

func (item ShijiesaiAll) ToList(withShengfu bool) ShijiesaiList {
	obj := ShijiesaiList{ID: item.ID}
	obj.No = item.No
	info := utils.StringToArr[XuanshouInfo](item.Info)
	if len(info) >= 1 {
		obj.AInfo = info[0]
	}
	if len(info) >= 2 {
		obj.BInfo = info[1]
	}
	if withShengfu {
		obj.ShengfuList = utils.StringToArr[int](item.ShengfuList)
	}
	return obj
}

type ShijiesaiAddData struct {
	Data []ShijiesaiList `json:"data"`
}
