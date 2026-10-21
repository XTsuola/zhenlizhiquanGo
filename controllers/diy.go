package controllers

import (
	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

// ---------- 创意皮肤 ----------

// skinDiyList 创意皮肤列表
func skinDiyList(c *gin.Context) {
	data, ok := findAll[models.SkinDiyAll]("skin_diy", c)
	if !ok {
		return
	}
	SearchList("查询成功", c, data)
}

// skinDiyAdd 新增创意皮肤
func skinDiyAdd(c *gin.Context) {
	params, ok := bindJSON[models.SkinDiyBase](c)
	if !ok {
		return
	}
	if err := my.DB.Table("skin_diy").Create(&params).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

// skinDiyAddAll 批量新增创意皮肤
func skinDiyAddAll(c *gin.Context) {
	params, ok := bindJSON[models.BatchData[models.SkinDiyBase]](c)
	if !ok {
		return
	}
	if !createBatch("skin_diy", params.Data, c) {
		return
	}
	HandleOk(c, "新增成功")
}

// skinDiyFields 创意皮肤可更新字段
func skinDiyFields(params models.SkinDiyUpdate) map[string]interface{} {
	return map[string]interface{}{
		"cardId": params.CardId,
		"name":   params.Name,
		"skill":  params.Skill,
		"effect": params.Effect,
		"reason": params.Reason,
		"remark": params.Remark,
	}
}

// skinDiyUpdate 管理员密码修改创意皮肤
func skinDiyUpdate(c *gin.Context) {
	params, ok := bindJSON[models.SkinDiyUpdate](c)
	if !ok {
		return
	}
	if !requireAdminPassword(params.Password, c) {
		return
	}
	if !updateByID("skin_diy", params.ID, skinDiyFields(params), c) {
		return
	}
	HandleOk(c, "操作成功")
}

// skinDiyUpdateTemp 临时密码修改创意皮肤（用后即删）
func skinDiyUpdateTemp(c *gin.Context) {
	params, ok := bindJSON[models.SkinDiyUpdate](c)
	if !ok {
		return
	}
	pwdID, ok := consumeTempPassword(params.Password, c)
	if !ok {
		return
	}
	if !updateByID("skin_diy", params.ID, skinDiyFields(params), c) {
		return
	}
	if !deleteTempPassword(pwdID, c) {
		return
	}
	HandleOk(c, "操作成功")
}

// ---------- 创意卡牌 ----------

// cardDiyList 创意卡牌列表
func cardDiyList(c *gin.Context) {
	data, ok := findAll[models.CardDiyAll]("card_diy", c)
	if !ok {
		return
	}
	SearchList("查询成功", c, data)
}

// cardDiyAdd 新增创意卡牌
func cardDiyAdd(c *gin.Context) {
	params, ok := bindJSON[models.CardDiyBase](c)
	if !ok {
		return
	}
	if err := my.DB.Table("card_diy").Create(&params).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

// cardDiyAddAll 批量新增创意卡牌
func cardDiyAddAll(c *gin.Context) {
	params, ok := bindJSON[models.BatchData[models.CardDiyBase]](c)
	if !ok {
		return
	}
	if !createBatch("card_diy", params.Data, c) {
		return
	}
	HandleOk(c, "新增成功")
}

// cardDiyFields 创意卡牌可更新字段
func cardDiyFields(params models.CardDiyUpdate) map[string]interface{} {
	return map[string]interface{}{
		"name":     params.Name,
		"zhenyin":  params.Zhenyin,
		"cost":     params.Cost,
		"quality":  params.Quality,
		"cardType": params.CardType,
		"att":      params.Att,
		"life":     params.Life,
		"effect":   params.Effect,
		"info":     params.Info,
		"remark":   params.Remark,
	}
}

// cardDiyUpdate 管理员密码修改创意卡牌
func cardDiyUpdate(c *gin.Context) {
	params, ok := bindJSON[models.CardDiyUpdate](c)
	if !ok {
		return
	}
	if !requireAdminPassword(params.Password, c) {
		return
	}
	if !updateByID("card_diy", params.ID, cardDiyFields(params), c) {
		return
	}
	HandleOk(c, "操作成功")
}

// cardDiyUpdateTemp 临时密码修改创意卡牌（用后即删）
func cardDiyUpdateTemp(c *gin.Context) {
	params, ok := bindJSON[models.CardDiyUpdate](c)
	if !ok {
		return
	}
	pwdID, ok := consumeTempPassword(params.Password, c)
	if !ok {
		return
	}
	if !updateByID("card_diy", params.ID, cardDiyFields(params), c) {
		return
	}
	if !deleteTempPassword(pwdID, c) {
		return
	}
	HandleOk(c, "操作成功")
}
