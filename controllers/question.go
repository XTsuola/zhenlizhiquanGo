package controllers

import (
	my "go_project/config"
	"go_project/models"
	"time"

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

func questionAddAll(c *gin.Context) {
	var params models.QuestionAddData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.QuestionBase
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.Info = item.Info
		data.Time = item.Time
		result := my.DB.Table("question").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}

func answerAdd(c *gin.Context) {
	var params models.AnswerBase
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("answer").Create(&params)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func answerList(c *gin.Context) {
	questionId := c.Param("questionId")
	var data []models.AnswerAll
	result := my.DB.Table("answer").Where("questionId = ?", questionId).Order("time DESC").Find(&data)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.AnswerAll]("查询成功", c, data)
}

func answerAddAll(c *gin.Context) {
	var params models.AnswerAddData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.AnswerBase
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.QuestionId = item.QuestionId
		data.Name = item.Name
		data.Content = item.Content
		data.Time = item.Time
		result := my.DB.Table("answer").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}

func answerAllList(c *gin.Context) {
	var data []models.AnswerAll
	result := my.DB.Table("answer").Order("time DESC").Find(&data)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.AnswerAll]("查询成功", c, data)
}
