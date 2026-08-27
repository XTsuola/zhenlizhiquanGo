package controllers

import (
	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

// 题目列表
func questionList(c *gin.Context) {
	data, ok := findAll[models.QuestionAll]("question", c)
	if !ok {
		return
	}
	SearchList("查询成功", c, data)
}

// 新增题目
func questionAdd(c *gin.Context) {
	params, ok := bindJSON[models.QuestionBase](c)
	if !ok {
		return
	}
	if err := my.DB.Table("question").Create(&params).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

// 题目详情
func questionDetail(c *gin.Context) {
	var data models.QuestionAll
	if err := my.DB.Table("question").Order("id desc").First(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchOne("查询成功", c, data)
}

// 批量新增题目
func questionAddAll(c *gin.Context) {
	params, ok := bindJSON[models.QuestionAddData](c)
	if !ok {
		return
	}
	if !createBatch("question", params.Data, c) {
		return
	}
	HandleOk(c, "新增成功")
}

// 新增答案
func answerAdd(c *gin.Context) {
	params, ok := bindJSON[models.AnswerBase](c)
	if !ok {
		return
	}
	if err := my.DB.Table("answer").Create(&params).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

// 问题对应答案列表
func answerList(c *gin.Context) {
	questionId := queryInt(c, "questionId")
	var data []models.AnswerAll
	if err := my.DB.Table("answer").Where("questionId = ?", questionId).Order("time DESC").Find(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchList("查询成功", c, data)
}

// 批量新增答案
func answerAddAll(c *gin.Context) {
	params, ok := bindJSON[models.AnswerAddData](c)
	if !ok {
		return
	}
	if !createBatch("answer", params.Data, c) {
		return
	}
	HandleOk(c, "新增成功")
}

// 所有答案列表
func answerAllList(c *gin.Context) {
	var data []models.AnswerAll
	if err := my.DB.Table("answer").Order("time DESC").Find(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchList("查询成功", c, data)
}
