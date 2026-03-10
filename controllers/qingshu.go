package controllers

import (
	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
	"math/rand"
	"time"
)

// 获取数据
func qingshuGetMap(c *gin.Context) {
	var mapObj models.QingshuMapParams
	result := my.DB.Table("qingshu").Where("id = ?", 1).First(&mapObj)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	var data models.QingshuMapData
	data.ID = mapObj.ID
	data.Round = mapObj.Round
	data.Status = mapObj.Status
	data.Msg = mapObj.Msg
	data.CardPile = StringToArr[int](mapObj.CardPile)
	data.DisPile = StringToArr[int](mapObj.DisPile)
	data.UserData = StringToArr[models.QingshuUserData](mapObj.UserData)
	SearchOne("查询成功", c, data)
}

// 重置游戏
func qingshuReset(c *gin.Context) {
	//cardBaseList := []int{3, 2, 6, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7, 8}
	cardBaseList := []int{1, 1, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 7, 8}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cardBaseList), func(i, j int) {
		cardBaseList[i], cardBaseList[j] = cardBaseList[j], cardBaseList[i]
	})
	var params models.QingshuMapData
	params.CardPile = cardBaseList[2:13]
	params.DisPile = cardBaseList[13:16]
	var user1 models.QingshuUserData
	user1.ID = 1
	user1.UserName = "用户1"
	user1.HandCards = cardBaseList[0:1]
	user1.DisCards = make([]int, 0)
	user1.Status = 1
	var user2 models.QingshuUserData
	user2.ID = 2
	user2.UserName = "用户2"
	user2.HandCards = cardBaseList[1:2]
	user2.DisCards = make([]int, 0)
	user2.Status = 1
	params.UserData = append(params.UserData, user1)
	params.UserData = append(params.UserData, user2)
	params.Round = 1
	params.Status = 0
	params.Msg = ""
	result := my.DB.Table("qingshu").Where("id = ?", 1).Updates(map[string]interface{}{
		"round":    params.Round,
		"status":   params.Status,
		"msg":      params.Msg,
		"cardPile": ArrToString(params.CardPile),
		"disPile":  ArrToString(params.DisPile),
		"userData": ArrToString(params.UserData),
	})
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "重置成功")
}

// 修改用户昵称
func userNameUpdate(c *gin.Context) {
	var params models.UsernameUpdate
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var mapObj models.QingshuMapParams
	result := my.DB.Table("qingshu").Where("id = ?", 1).First(&mapObj)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	var data models.QingshuMapData
	data.UserData = StringToArr[models.QingshuUserData](mapObj.UserData)
	if params.Password == "1" {
		data.UserData[0].UserName = params.Name
	} else if params.Password == "2" {
		data.UserData[1].UserName = params.Name
	}
	result2 := my.DB.Table("qingshu").Where("id = ?", 1).Updates(map[string]interface{}{
		"userData": ArrToString(data.UserData),
	})
	if result2.Error != nil {
		MyErr(result2.Error.Error(), c)
		return
	}
	HandleOk(c, "操作成功")
}
