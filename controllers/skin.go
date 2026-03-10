package controllers

import (
	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
	"time"
)

func skinList(c *gin.Context) {
	var list []models.SkinSelect
	result := my.DB.Table("skin").Find(&list)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	var data []models.SkinAll
	for _, item := range list {
		var obj models.SkinAll
		obj.ID = item.ID
		obj.CardId = item.CardId
		obj.Name = item.Name
		obj.Zhenyin = item.Zhenyin
		obj.Cost = item.Cost
		obj.Skill = item.Skill
		obj.Img = item.Img
		obj.Shuxing = item.Shuxing
		obj.Origin = item.Origin
		obj.Remark = item.Remark
		obj.Together = item.Together
		obj.Effect = StringToArr[string](item.Effect)
		data = append(data, obj)
	}
	SearchList[models.SkinAll]("查询成功", c, data)
}

func skinTogether(c *gin.Context) {
	var params models.SkinTogetherParams
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("skin").Where("id = ?", params.ID).Updates(map[string]interface{}{
		"together": params.Together,
	})
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "操作成功")
}

func skinAdd(c *gin.Context) {
	var params models.SkinAddData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.SkinAddObj
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.CardId = item.CardId
		data.Name = item.Name
		data.Zhenyin = item.Zhenyin
		data.Cost = item.Cost
		data.Skill = item.Skill
		data.Img = item.Img
		data.Shuxing = item.Shuxing
		data.Origin = item.Origin
		data.Effect = ArrToString(item.Effect)
		data.Remark = item.Remark
		data.Together = item.Together
		result := my.DB.Table("skin").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}

func skinTogetherAll(c *gin.Context) {
	var params models.SkinUpdateData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		result := my.DB.Table("skin").Where("id = ?", item.ID).Updates(map[string]interface{}{
			"together": item.Together,
		})
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "操作成功")
}
