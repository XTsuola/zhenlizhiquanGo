package controllers

import (
	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

// 卡组列字段
func frequencyFields(params models.FrequencyUpdate) map[string]interface{} {
	return map[string]interface{}{
		"name":     params.Name,
		"qu":       params.Qu,
		"heroId":   params.HeroId,
		"heroLife": params.HeroLife,
		"cards":    params.Cards,
		"time":     params.Time,
	}
}

// 新增卡组
func frequencyCardsAdd(c *gin.Context) {
	params, ok := bindJSON[models.FrequencyBase](c)
	if !ok {
		return
	}
	if err := my.DB.Table("frequency").Create(&params).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

// 修改卡组
func frequencyCardsUpdate(c *gin.Context) {
	params, ok := bindJSON[models.FrequencyUpdate](c)
	if !ok {
		return
	}
	if !requireAdminPassword(params.Password, c) {
		return
	}
	if !updateByID("frequency", params.ID, frequencyFields(params), c) {
		return
	}
	HandleOk(c, "操作成功")
}

// 非管理员临时修改卡组
func frequencyCardsUpdateTemp(c *gin.Context) {
	params, ok := bindJSON[models.FrequencyUpdate](c)
	if !ok {
		return
	}
	pwdID, ok := consumeTempPassword(params.Password, c)
	if !ok {
		return
	}
	if !updateByID("frequency", params.ID, frequencyFields(params), c) {
		return
	}
	if !deleteTempPassword(pwdID, c) {
		return
	}
	HandleOk(c, "操作成功")
}

// 卡组详情
func frequencyCardsDetail(c *gin.Context) {
	id := queryInt(c, "id")
	var list []models.FrequencyAll
	if err := my.DB.Table("frequency").Where("heroId = ?", id).Find(&list).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchList("查询成功", c, list)
}

// 卡组列表
func frequencyCardsAll(c *gin.Context) {
	list, ok := findAll[models.FrequencyAll]("frequency", c)
	if !ok {
		return
	}
	SearchList("查询成功", c, list)
}

// 删除卡组
func frequencyCardsDelete(c *gin.Context) {
	if !requireAdminPassword(c.Query("password"), c) {
		return
	}
	if !deleteByQuery("frequency", c, "id = ?", queryInt(c, "id")) {
		return
	}
	HandleOk(c, "删除成功")
}

// 批量新增卡组
func frequencyCardsAddAll(c *gin.Context) {
	params, ok := bindJSON[models.FrequencyAddAll](c)
	if !ok {
		return
	}
	if !createBatch("frequency", params.Data, c) {
		return
	}
	HandleOk(c, "新增成功")
}

// 新增编辑卡组密码
func frequencyPasswordAdd(c *gin.Context) {
	params, ok := bindJSON[models.PasswordAdd](c)
	if !ok {
		return
	}
	if err := my.DB.Table("password").Create(&params).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

// 修改编辑卡组密码
func frequencyPasswordList(c *gin.Context) {
	list, ok := findAll[models.PasswordAll]("password", c)
	if !ok {
		return
	}
	SearchList("查询成功", c, list)
}

// 删除编辑卡组密码
func frequencyPasswordDelete(c *gin.Context) {
	if !deleteByQuery("password", c, "id = ?", queryInt(c, "id")) {
		return
	}
	HandleOk(c, "删除成功")
}
