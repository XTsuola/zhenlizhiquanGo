package controllers

import (
	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
	"strconv"
	"time"
)

func shenqiList(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("zhenyin"))
	var list []models.ShenqiSelect
	result := my.DB.Table("shenqi").Where("zhenyin = ?", id).Find(&list)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	var data []models.ShenqiAll
	for _, item := range list {
		var obj models.ShenqiAll
		obj.ID = item.ID
		obj.Name = item.Name
		obj.Zhenyin = item.Zhenyin
		obj.Quality = item.Quality
		obj.Type = item.Type
		obj.Img = item.Img
		obj.Data = StringToArr[models.ShenqiData](item.Data)
		data = append(data, obj)
	}
	SearchList[models.ShenqiAll]("查询成功", c, data)
}

func shenqiAdd(c *gin.Context) {
	var params models.ShenqiAddData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.ShenqiAddObj
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.Name = item.Name
		data.Zhenyin = item.Zhenyin
		data.Quality = item.Quality
		data.Type = item.Type
		data.Img = item.Img
		data.Data = ArrToString[models.ShenqiData](item.Data)
		result := my.DB.Table("shenqi").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}
