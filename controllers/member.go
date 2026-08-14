package controllers

import (
	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

func memberList(c *gin.Context) {
	var data []models.MemberAll
	if err := my.DB.Table("member").Order("donation desc, score desc").Find(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchList("查询成功", c, data)
}

func memberAdd(c *gin.Context) {
	params, ok := bindJSON[models.MemberBase](c)
	if !ok {
		return
	}
	if err := my.DB.Table("member").Create(&params).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func memberUpdate(c *gin.Context) {
	params, ok := bindJSON[models.MemberBase](c)
	if !ok {
		return
	}
	if !updateByID("member", c.Param("id"), map[string]interface{}{
		"name":     params.Name,
		"donation": params.Donation,
		"reward":   params.Reward,
		"score":    params.Score,
		"title":    params.Title,
		"remark":   params.Remark,
	}, c) {
		return
	}
	HandleOk(c, "操作成功")
}

func memberDelete(c *gin.Context) {
	if !deleteByQuery("member", c, "id = ?", c.Param("id")) {
		return
	}
	HandleOk(c, "删除成功")
}

func memberRewardList(c *gin.Context) {
	var data []models.MemberAll
	if err := my.DB.Table("member").Where("reward > ?", 0).Order("reward desc").Find(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchList("查询成功", c, data)
}

func memberAddAll(c *gin.Context) {
	params, ok := bindJSON[models.MemberAddData](c)
	if !ok {
		return
	}
	if !createBatch("member", params.Data, c) {
		return
	}
	HandleOk(c, "新增成功")
}
