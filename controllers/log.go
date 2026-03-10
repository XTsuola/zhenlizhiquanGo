package controllers

import (
	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
	"time"
)

func logList(c *gin.Context) {
	var data []models.LogList
	result := my.DB.Table("log").Order("time DESC").Find(&data)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.LogList]("查询成功", c, data)
}

func logAdd(c *gin.Context) {
	name := c.Query("name")
	var data models.LogAdd
	data.Name = name
	data.Time = time.Now().Format("2006-01-02 15:04:05")
	result := my.DB.Table("log").Create(&data)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}
