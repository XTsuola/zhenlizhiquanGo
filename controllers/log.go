package controllers

import (
	"time"

	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

// 日志列表
func logList(c *gin.Context) {
	var data []models.LogList
	if err := my.DB.Table("log").Order("time DESC").Find(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchList("查询成功", c, data)
}

// 新增日志
func logAdd(c *gin.Context) {
	data := models.LogAdd{
		Name: c.Query("name"),
		Time: time.Now().Format("2006-01-02 15:04:05"),
	}
	if err := my.DB.Table("log").Create(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}
