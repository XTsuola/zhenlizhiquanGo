package controllers

import (
	"go_project/models"

	my "go_project/config"
	"github.com/gin-gonic/gin"
)

func noteList(c *gin.Context) {
	data, ok := findAll[models.NoteAll]("note", c)
	if !ok {
		return
	}
	SearchList("查询成功", c, data)
}

func noteAdd(c *gin.Context) {
	params, ok := bindJSON[models.NoteBase](c)
	if !ok {
		return
	}
	if err := my.DB.Table("note").Create(&params).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func noteDelete(c *gin.Context) {
	if !deleteByQuery("note", c, "id = ?", queryInt(c, "id")) {
		return
	}
	HandleOk(c, "删除成功")
}
