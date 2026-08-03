package controllers

import (
	"go_project/models"

	"github.com/gin-gonic/gin"
)

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

func shardList(c *gin.Context) {
	list, ok := findAll[models.ShardAll]("shard", c)
	if !ok {
		return
	}
	SearchList("查询成功", c, list)
}

func shardUpdate(c *gin.Context) {
	params, ok := bindJSON[models.ShardUpdateParams](c)
	if !ok {
		return
	}
	if !updateByID("shard", params.ID, map[string]interface{}{
		"skillData": params.SkillData,
	}, c) {
		return
	}
	HandleOk(c, "操作成功")
}

func shardAdd(c *gin.Context) {
	params, ok := bindJSON[models.ShardAddData](c)
	if !ok {
		return
	}
	if !createBatch("shard", params.Data, c) {
		return
	}
	HandleOk(c, "新增成功")
}
