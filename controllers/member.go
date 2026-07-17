package controllers

import (
	my "go_project/config"
	"go_project/models"
	"time"

	"github.com/gin-gonic/gin"
)

func memberList(c *gin.Context) {
	var data []models.MemberAll
	result := my.DB.Table("member").Order("donation desc, score desc").Find(&data)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.MemberAll]("查询成功", c, data)
}

func memberAdd(c *gin.Context) {
	var params models.MemberBase
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("member").Create(&params)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func memberUpdate(c *gin.Context) {
	id := c.Param("id")
	var params models.MemberBase
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("member").Where("id = ?", id).Updates(map[string]interface{}{
		"name":     params.Name,
		"donation": params.Donation,
		"score":    params.Score,
		"title":    params.Title,
		"remark":   params.Remark,
	})
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "操作成功")
}

func memberDelete(c *gin.Context) {
	id := c.Param("id")
	result := my.DB.Table("member").Where("id = ?", id).Delete(nil)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "删除成功")
}

func memberAddAll(c *gin.Context) {
	var params models.MemberAddData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.MemberBase
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.Name = item.Name
		data.Donation = item.Donation
		data.Score = item.Score
		data.Title = item.Title
		data.Remark = item.Remark
		result := my.DB.Table("member").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}
