package controllers

import (
	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

// 神器列表
func shenqiList(c *gin.Context) {
	id := queryInt(c, "zhenyin")
	var list []models.ShenqiSelect
	if err := my.DB.Table("shenqi").Where("zhenyin = ?", id).Find(&list).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	data := make([]models.ShenqiAll, 0, len(list))
	for _, item := range list {
		data = append(data, item.ToAll())
	}
	SearchList("查询成功", c, data)
}

// 批量新增神器
func shenqiAdd(c *gin.Context) {
	params, ok := bindJSON[models.ShenqiAddData](c)
	if !ok {
		return
	}
	rows := make([]models.ShenqiAddObj, 0, len(params.Data))
	for _, item := range params.Data {
		rows = append(rows, item.ToObj())
	}
	if !createBatch("shenqi", rows, c) {
		return
	}
	HandleOk(c, "新增成功")
}
