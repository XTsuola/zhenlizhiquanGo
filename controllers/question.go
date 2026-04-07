package controllers

import (
	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

func questionList(c *gin.Context) {
	var data []models.QuestionAll
	result := my.DB.Table("question").Find(&data)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.QuestionAll]("查询成功", c, data)
}

func questionAdd(c *gin.Context) {
	var params models.QuestionBase
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("question").Create(&params)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func questionDetail(c *gin.Context) {
	var data models.QuestionAll
	_ = my.DB.Table("question").Order("id desc").First(&data)
	SearchOne[models.QuestionAll]("查询成功", c, data)
}

func questionAnswer(c *gin.Context) {

}

//func noteDelete(c *gin.Context) {
//	id, _ := strconv.Atoi(c.Query("id"))
//	result := my.DB.Table("note").Where("id = ?", id).Delete(nil)
//	if result.Error != nil {
//		MyErr(result.Error.Error(), c)
//		return
//	}
//	HandleOk(c, "删除成功")
//}
