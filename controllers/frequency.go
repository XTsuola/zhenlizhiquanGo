package controllers

import (
	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

// frequencyFields 组装卡组可更新字段，不含密码
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

// frequencyCardsAdd 新增卡组
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

// frequencyCardsUpdate 管理员密码修改卡组
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

// frequencyCardsUpdateTemp 临时密码修改卡组，成功后删除该密码
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

// frequencyCardsDetail 按英雄 ID 查询卡组
func frequencyCardsDetail(c *gin.Context) {
	id := queryInt(c, "id")
	var list []models.FrequencyAll
	if err := my.DB.Table("frequency").Where("heroId = ?", id).Find(&list).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchList("查询成功", c, list)
}

// frequencyCardsAll 全部卡组
func frequencyCardsAll(c *gin.Context) {
	list, ok := findAll[models.FrequencyAll]("frequency", c)
	if !ok {
		return
	}
	SearchList("查询成功", c, list)
}

// frequencyCardsDelete 管理员密码按 ID 删除卡组
func frequencyCardsDelete(c *gin.Context) {
	if !requireAdminPassword(c.Query("password"), c) {
		return
	}
	if !deleteByQuery("frequency", c, "id = ?", queryInt(c, "id")) {
		return
	}
	HandleOk(c, "删除成功")
}

// frequencyCardsAddAll 批量新增卡组
func frequencyCardsAddAll(c *gin.Context) {
	params, ok := bindJSON[models.BatchData[models.FrequencyBase]](c)
	if !ok {
		return
	}
	if !createBatch("frequency", params.Data, c) {
		return
	}
	HandleOk(c, "新增成功")
}

// frequencyPasswordAdd 新增一次性临时密码
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

// frequencyPasswordList 临时密码列表
func frequencyPasswordList(c *gin.Context) {
	list, ok := findAll[models.PasswordAll]("password", c)
	if !ok {
		return
	}
	SearchList("查询成功", c, list)
}

// frequencyPasswordDelete 按 ID 删除临时密码
func frequencyPasswordDelete(c *gin.Context) {
	if !deleteByQuery("password", c, "id = ?", queryInt(c, "id")) {
		return
	}
	HandleOk(c, "删除成功")
}
