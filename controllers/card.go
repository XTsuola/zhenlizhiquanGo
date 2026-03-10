package controllers

import (
	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
	"strconv"
	"time"
)

func cardList(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("zhenyin"))
	var list []models.CardSelect
	result := my.DB.Table("card").Where("zhenyin = ?", id).Find(&list)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	var data []models.CardAll
	for _, item := range list {
		var obj models.CardAll
		obj.ID = item.ID
		obj.Name = item.Name
		obj.Zhenyin = item.Zhenyin
		obj.Quality = item.Quality
		obj.Cost = item.Cost
		obj.Type = item.Type
		obj.Img = item.Img
		obj.Grade = item.Grade
		obj.Tag = item.Tag
		obj.Data = StringToArr[models.CardData](item.Data)
		data = append(data, obj)
	}
	SearchList[models.CardAll]("查询成功", c, data)
}

func cardAllList(c *gin.Context) {
	var list []models.CardSelect
	result := my.DB.Table("card").Find(&list)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	var data []models.CardAll
	for _, item := range list {
		var obj models.CardAll
		obj.ID = item.ID
		obj.Name = item.Name
		obj.Zhenyin = item.Zhenyin
		obj.Quality = item.Quality
		obj.Cost = item.Cost
		obj.Type = item.Type
		obj.Img = item.Img
		obj.Grade = item.Grade
		obj.Tag = item.Tag
		obj.Data = StringToArr[models.CardData](item.Data)
		data = append(data, obj)
	}
	SearchList[models.CardAll]("查询成功", c, data)
}

func cardAdd(c *gin.Context) {
	var params models.CardAddData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.CardAddObj
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.Name = item.Name
		data.Zhenyin = item.Zhenyin
		data.Quality = item.Quality
		data.Cost = item.Cost
		data.Type = item.Type
		data.Img = item.Img
		data.Grade = item.Grade
		data.Tag = item.Tag
		data.Data = ArrToString[models.CardData](item.Data)
		result := my.DB.Table("card").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}

func cardGradeUpdate(c *gin.Context) {
	var params models.CardUpdateGradeParams
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("card").Where("id = ?", params.ID).Updates(map[string]interface{}{
		"grade": ArrToString(params.Grade),
	})
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "操作成功")
}

func cardGradeUpdateList(c *gin.Context) {
	var params models.CardUpdateGradeListData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		result := my.DB.Table("card").Where("id = ?", item.ID).Updates(map[string]interface{}{
			"grade": item.Grade,
		})
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}

func cardTagUpdate(c *gin.Context) {
	var params models.CardUpdateTagParams
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.CardUpdateTag
	data.ID = params.ID
	data.Tag = ArrToString(params.Tag)
	result := my.DB.Table("card").Where("id = ?", params.ID).Updates(map[string]interface{}{
		"tag": data.Tag,
	})
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "操作成功")
}

func cardTagUpdateList(c *gin.Context) {
	var params models.CardUpdateTagListData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		result := my.DB.Table("card").Where("id = ?", item.ID).Updates(map[string]interface{}{
			"tag": item.Tag,
		})
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}
