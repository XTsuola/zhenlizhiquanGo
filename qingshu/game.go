package qingshu

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	my "go_project/config"
	"go_project/models"
	"go_project/utils"

	"gorm.io/gorm"
)

// NewGame 生成一局初始牌局；可保留玩家昵称
func NewGame(user1, user2 string) models.QingshuMapData {
	if user1 == "" {
		user1 = "用户1"
	}
	if user2 == "" {
		user2 = "用户2"
	}
	deck := []int{1, 1, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 7, 8}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	return models.QingshuMapData{
		ID:       1,
		CardPile: deck[2:13],
		DisPile:  deck[13:16],
		UserData: []models.QingshuUserData{
			{ID: 1, UserName: user1, HandCards: append([]int{}, deck[0]), DisCards: []int{}, Status: 1},
			{ID: 2, UserName: user2, HandCards: append([]int{}, deck[1]), DisCards: []int{}, Status: 1},
		},
		QingshuMapBase: models.QingshuMapBase{Round: 1, Status: 0, Msg: ""},
	}
}

func toRow(data models.QingshuMapData) models.QingshuMapParams {
	return models.QingshuMapParams{
		ID:             1,
		CardPile:       utils.ArrToString(data.CardPile),
		DisPile:        utils.ArrToString(data.DisPile),
		UserData:       utils.ArrToString(data.UserData),
		QingshuMapBase: data.QingshuMapBase,
	}
}

// LoadOrInit 读取 id=1，不存在则自动建局
func LoadOrInit() (models.QingshuMapParams, error) {
	var row models.QingshuMapParams
	err := my.DB.Table("qingshu").Where("id = ?", 1).First(&row).Error
	if err == nil {
		return row, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return row, err
	}
	row = toRow(NewGame("", ""))
	if err = my.DB.Table("qingshu").Create(&row).Error; err != nil {
		return row, err
	}
	return row, nil
}

// Save 更新现有对局
func Save(data models.QingshuMapData) error {
	return my.DB.Table("qingshu").Where("id = ?", 1).Updates(data.ToUpdateMap()).Error
}

// SaveOrCreate Updates；无行则 Create
func SaveOrCreate(data models.QingshuMapData) error {
	result := my.DB.Table("qingshu").Where("id = ?", 1).Updates(data.ToUpdateMap())
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return my.DB.Table("qingshu").Create(toRow(data)).Error
	}
	return nil
}

// Reset 重置牌局；roundStarter>0 时作为先手 round（WS 兼容）
func Reset(roundStarter int) error {
	row, err := LoadOrInit()
	if err != nil {
		return err
	}
	data := row.ToData()
	u1, u2 := "用户1", "用户2"
	if len(data.UserData) >= 1 && data.UserData[0].UserName != "" {
		u1 = data.UserData[0].UserName
	}
	if len(data.UserData) >= 2 && data.UserData[1].UserName != "" {
		u2 = data.UserData[1].UserName
	}
	next := NewGame(u1, u2)
	if roundStarter > 0 {
		next.Round = roundStarter
	}
	return SaveOrCreate(next)
}

// Draw 摸牌
func Draw(userId int) error {
	row, err := LoadOrInit()
	if err != nil {
		return err
	}
	data := row.ToData()
	if len(data.CardPile) == 0 {
		return fmt.Errorf("牌堆已空")
	}
	if len(data.UserData) < 2 {
		return fmt.Errorf("用户数据异常")
	}
	card := data.CardPile[0]
	data.CardPile = data.CardPile[1:]
	data.Status = 1
	data.Msg = ""
	if userId == 1 {
		data.UserData[0].HandCards = append(data.UserData[0].HandCards, card)
		data.UserData[0].Status = 1
	} else {
		data.UserData[1].HandCards = append(data.UserData[1].HandCards, card)
		data.UserData[1].Status = 1
	}
	return Save(data)
}

// Play 出牌；返回更新后的牌局
func Play(myId, pai, yourPai, index int) (models.QingshuMapData, error) {
	row, err := LoadOrInit()
	if err != nil {
		return models.QingshuMapData{}, err
	}
	obj := row.ToData()
	if len(obj.UserData) < 2 {
		return models.QingshuMapData{}, fmt.Errorf("用户数据异常")
	}

	params := models.QingshuMapData{
		CardPile: append([]int{}, obj.CardPile...),
		DisPile:  append([]int{}, obj.DisPile...),
		QingshuMapBase: models.QingshuMapBase{
			Round:  obj.Round + 1,
			Status: obj.Status,
			Msg:    "",
		},
	}

	var me, you models.QingshuUserData
	if myId == 1 {
		me, you = obj.UserData[0], obj.UserData[1]
	} else {
		me, you = obj.UserData[1], obj.UserData[0]
	}
	me.HandCards = append([]int{}, me.HandCards...)
	me.DisCards = append([]int{}, me.DisCards...)
	you.HandCards = append([]int{}, you.HandCards...)
	you.DisCards = append([]int{}, you.DisCards...)

	needHand := 1
	if index == 0 {
		needHand = 2
	}
	if len(me.HandCards) < needHand {
		return models.QingshuMapData{}, fmt.Errorf("手牌不足")
	}

	if you.Status != 3 {
		switch pai {
		case 1:
			if len(you.HandCards) == 0 {
				return models.QingshuMapData{}, fmt.Errorf("对方无手牌")
			}
			if yourPai == you.HandCards[0] {
				you.Status = 2
				you.DisCards = append(you.DisCards, you.HandCards[0])
				you.HandCards = []int{}
				params.Status = 2
			}
		case 2:
			if len(you.HandCards) == 0 {
				return models.QingshuMapData{}, fmt.Errorf("对方无手牌")
			}
			params.Msg = strconv.Itoa(you.HandCards[0])
		case 3:
			if len(you.HandCards) == 0 {
				return models.QingshuMapData{}, fmt.Errorf("对方无手牌")
			}
			myCard := me.HandCards[0]
			if index == 0 {
				myCard = me.HandCards[1]
			}
			if myCard > you.HandCards[0] {
				you.Status = 2
				params.Status = 2
			} else if myCard < you.HandCards[0] {
				me.Status = 2
				params.Status = 2
			}
		case 5:
			if len(you.HandCards) == 0 {
				return models.QingshuMapData{}, fmt.Errorf("对方无手牌")
			}
			you.DisCards = append(you.DisCards, you.HandCards[0])
			if you.HandCards[0] == 8 {
				you.Status = 2
				you.HandCards = []int{}
				params.Status = 2
			} else if len(params.CardPile) > 0 {
				you.HandCards[0] = params.CardPile[0]
				params.CardPile = params.CardPile[1:]
			} else if len(params.DisPile) > 0 {
				you.HandCards[0] = params.DisPile[0]
				params.DisPile = params.DisPile[1:]
			} else {
				return models.QingshuMapData{}, fmt.Errorf("无牌可补")
			}
		case 6:
			if len(you.HandCards) == 0 {
				return models.QingshuMapData{}, fmt.Errorf("对方无手牌")
			}
			youCard := you.HandCards[0]
			if index == 0 {
				you.HandCards[0] = me.HandCards[1]
				me.HandCards[1] = youCard
			} else {
				you.HandCards[0] = me.HandCards[0]
				me.HandCards[0] = youCard
			}
		}
	}

	switch pai {
	case 4:
		me.Status = 3
	case 8:
		me.Status = 2
		params.Status = 2
	}

	if index == 0 {
		me.HandCards = me.HandCards[1:]
	} else {
		me.HandCards = me.HandCards[:1]
	}
	me.DisCards = append(me.DisCards, pai)

	if len(params.CardPile) == 0 {
		if len(me.HandCards) == 0 || len(you.HandCards) == 0 {
			return models.QingshuMapData{}, fmt.Errorf("结算时手牌异常")
		}
		if me.HandCards[0] > you.HandCards[0] {
			you.Status = 2
			params.Status = 2
		} else if me.HandCards[0] < you.HandCards[0] {
			me.Status = 2
			params.Status = 2
		} else {
			if len(params.DisPile) < 2 {
				return models.QingshuMapData{}, fmt.Errorf("弃牌堆不足")
			}
			me.HandCards[0] = params.DisPile[0]
			you.HandCards[0] = params.DisPile[1]
			params.DisPile = params.DisPile[2:]
			if me.HandCards[0] > you.HandCards[0] {
				you.Status = 2
				params.Status = 2
			} else {
				me.Status = 2
				params.Status = 2
			}
		}
	}

	if myId == 1 {
		params.UserData = []models.QingshuUserData{me, you}
	} else {
		params.UserData = []models.QingshuUserData{you, me}
	}
	if params.Status == 2 {
		params.Msg = "游戏结束"
	}
	if err = Save(params); err != nil {
		return models.QingshuMapData{}, err
	}
	return params, nil
}

// UpdateUsername password 为玩家槽位 "1"/"2"
func UpdateUsername(name, slot string) error {
	row, err := LoadOrInit()
	if err != nil {
		return err
	}
	data := row.ToData()
	if len(data.UserData) < 2 {
		return fmt.Errorf("用户数据异常")
	}
	switch slot {
	case "1":
		data.UserData[0].UserName = name
	case "2":
		data.UserData[1].UserName = name
	default:
		return fmt.Errorf("玩家标识错误，password 应为 1 或 2")
	}
	return my.DB.Table("qingshu").Where("id = ?", 1).Updates(map[string]interface{}{
		"userData": utils.ArrToString(data.UserData),
	}).Error
}
