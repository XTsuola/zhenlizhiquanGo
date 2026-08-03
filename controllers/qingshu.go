package controllers

import (
	"math/rand"
	"time"

	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

func qingshuGetMap(c *gin.Context) {
	var mapObj models.QingshuMapParams
	if err := my.DB.Table("qingshu").Where("id = ?", 1).First(&mapObj).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchOne("查询成功", c, mapObj.ToData())
}

func qingshuReset(c *gin.Context) {
	cardBaseList := []int{1, 1, 1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 7, 8}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(cardBaseList), func(i, j int) {
		cardBaseList[i], cardBaseList[j] = cardBaseList[j], cardBaseList[i]
	})

	params := models.QingshuMapData{
		CardPile: cardBaseList[2:13],
		DisPile:  cardBaseList[13:16],
		UserData: []models.QingshuUserData{
			{ID: 1, UserName: "用户1", HandCards: cardBaseList[0:1], DisCards: []int{}, Status: 1},
			{ID: 2, UserName: "用户2", HandCards: cardBaseList[1:2], DisCards: []int{}, Status: 1},
		},
		QingshuMapBase: models.QingshuMapBase{Round: 1, Status: 0, Msg: ""},
	}
	if err := my.DB.Table("qingshu").Where("id = ?", 1).Updates(params.ToUpdateMap()).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "重置成功")
}

func userNameUpdate(c *gin.Context) {
	params, ok := bindJSON[models.UsernameUpdate](c)
	if !ok {
		return
	}
	var mapObj models.QingshuMapParams
	if err := my.DB.Table("qingshu").Where("id = ?", 1).First(&mapObj).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	userData := StringToArr[models.QingshuUserData](mapObj.UserData)
	if len(userData) < 2 {
		MyErr("用户数据异常", c)
		return
	}
	switch params.Password {
	case "1":
		userData[0].UserName = params.Name
	case "2":
		userData[1].UserName = params.Name
	}
	if err := my.DB.Table("qingshu").Where("id = ?", 1).Updates(map[string]interface{}{
		"userData": ArrToString(userData),
	}).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "操作成功")
}
