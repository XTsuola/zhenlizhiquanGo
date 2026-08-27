package controllers

import (
	"go_project/models"

	"github.com/gin-gonic/gin"
)

// 英雄列表
func heroList(c *gin.Context) {
	list, ok := findAll[models.HeroSelect]("hero", c)
	if !ok {
		return
	}
	data := make([]models.HeroAll, 0, len(list))
	for _, item := range list {
		data = append(data, item.ToAll())
	}
	SearchList("查询成功", c, data)
}

// 批量新增英雄
func heroAdd(c *gin.Context) {
	params, ok := bindJSON[models.HeroAddData](c)
	if !ok {
		return
	}
	rows := make([]models.HeroAddObj, 0, len(params.Data))
	for _, item := range params.Data {
		rows = append(rows, item.ToObj())
	}
	if !createBatch("hero", rows, c) {
		return
	}
	HandleOk(c, "新增成功")
}

// 修改英雄代理人
func agentAdd(c *gin.Context) {
	params, ok := bindJSON[models.AgentUpdate](c)
	if !ok {
		return
	}
	if !updateByID("hero", params.ID, map[string]interface{}{
		"agent": ArrToString(params.Agent),
	}, c) {
		return
	}
	HandleOk(c, "操作成功")
}
