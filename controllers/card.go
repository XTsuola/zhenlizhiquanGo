package controllers

import (
	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

// 卡牌列表，根据种族查询：1.帝国、2.隐秘、3.禅意、4.港口、5.炼狱、6.蛮石、7.冬神
func cardList(c *gin.Context) {
	id := queryInt(c, "zhenyin")
	var list []models.CardSelect
	if err := my.DB.Table("card").Where("zhenyin = ?", id).Find(&list).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	data := make([]models.CardAll, 0, len(list))
	for _, item := range list {
		data = append(data, item.ToAll())
	}
	SearchList("查询成功", c, data)
}

// 所有卡牌列表，不分种族
func cardAllList(c *gin.Context) {
	list, ok := findAll[models.CardSelect]("card", c)
	if !ok {
		return
	}
	data := make([]models.CardAll, 0, len(list))
	for _, item := range list {
		data = append(data, item.ToAll())
	}
	SearchList("查询成功", c, data)
}

// 批量新增卡牌
func cardAdd(c *gin.Context) {
	params, ok := bindJSON[models.CardAddData](c)
	if !ok {
		return
	}
	rows := make([]models.CardAddObj, 0, len(params.Data))
	for _, item := range params.Data {
		rows = append(rows, item.ToObj())
	}
	if !createBatch("card", rows, c) {
		return
	}
	HandleOk(c, "新增成功")
}

// 修改卡牌评级
func cardGradeUpdate(c *gin.Context) {
	params, ok := bindJSON[models.CardUpdateGradeParams](c)
	if !ok {
		return
	}
	if !updateByID("card", params.ID, map[string]interface{}{
		"grade": ArrToString(params.Grade),
	}, c) {
		return
	}
	HandleOk(c, "操作成功")
}

// 修改卡牌标签
func cardTagUpdate(c *gin.Context) {
	params, ok := bindJSON[models.CardUpdateTagParams](c)
	if !ok {
		return
	}
	if !updateByID("card", params.ID, map[string]interface{}{
		"tag": ArrToString(params.Tag),
	}, c) {
		return
	}
	HandleOk(c, "操作成功")
}
