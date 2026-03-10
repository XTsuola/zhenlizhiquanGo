package controllers

import (
	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
	"strconv"
)

func noteList(c *gin.Context) {
	var data []models.NoteAll
	result := my.DB.Table("note").Find(&data)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.NoteAll]("查询成功", c, data)
}

func noteAdd(c *gin.Context) {
	var params models.NoteBase
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("note").Create(&params)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func noteDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	result := my.DB.Table("note").Where("id = ?", id).Delete(nil)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "删除成功")
}
