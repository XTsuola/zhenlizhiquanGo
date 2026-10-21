package controllers

import (
	"go_project/models"

	"github.com/gin-gonic/gin"
)

// heroList 英雄列表，代理人与技能数据从 JSON 还原
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

// heroAdd 批量新增英雄
func heroAdd(c *gin.Context) {
	params, ok := bindJSON[models.BatchData[models.HeroAddParams]](c)
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

// agentAdd 按 ID 更新英雄代理人
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
