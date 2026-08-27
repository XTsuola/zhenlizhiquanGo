package controllers

import (
	"go_project/models"

	"github.com/gin-gonic/gin"
)

// 皮肤列表
func skinList(c *gin.Context) {
	list, ok := findAll[models.SkinSelect]("skin", c)
	if !ok {
		return
	}
	data := make([]models.SkinAll, 0, len(list))
	for _, item := range list {
		data = append(data, item.ToAll())
	}
	SearchList("查询成功", c, data)
}

// 新增皮肤
func skinAdd(c *gin.Context) {
	params, ok := bindJSON[models.SkinAddData](c)
	if !ok {
		return
	}
	rows := make([]models.SkinAddObj, 0, len(params.Data))
	for _, item := range params.Data {
		rows = append(rows, item.ToObj())
	}
	if !createBatch("skin", rows, c) {
		return
	}
	HandleOk(c, "新增成功")
}
