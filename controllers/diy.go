package controllers

import (
	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
	"time"
)

func skinDiyList(c *gin.Context) {
	var data []models.SkinDiyAll
	result := my.DB.Table("skin_diy").Find(&data)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.SkinDiyAll]("查询成功", c, data)
}

func skinDiyAdd(c *gin.Context) {
	var params models.SkinDiyBase
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("skin_diy").Create(&params)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func skinDiyAddAll(c *gin.Context) {
	var params models.SkinDiyAddData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.SkinDiyBase
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.CardId = item.CardId
		data.Name = item.Name
		data.Skill = item.Skill
		data.Effect = item.Effect
		data.Reason = item.Reason
		data.Remark = item.Remark
		result := my.DB.Table("skin_diy").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}

func skinDiyUpdate(c *gin.Context) {
	var params models.SkinDiyUpdate
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	if params.Password != "suola18" {
		MyErr("管理员密码错误", c)
		return
	} else {
		result := my.DB.Table("skin_diy").Where("id = ?", params.ID).Updates(map[string]interface{}{
			"cardId": params.CardId,
			"name":   params.Name,
			"skill":  params.Skill,
			"effect": params.Effect,
			"reason": params.Reason,
			"remark": params.Remark,
		})
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
		HandleOk(c, "操作成功")
	}
}

func skinDiyUpdateTemp(c *gin.Context) {
	var params models.SkinDiyUpdate
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var obj models.FrequencyPaddwordAll
	result := my.DB.Table("password").Where("password = ?", params.Password).Find(&obj)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	if result.RowsAffected == 0 {
		MyErr("临时密码错误", c)
		return
	} else {
		result2 := my.DB.Table("skin_diy").Where("id = ?", params.ID).Updates(map[string]interface{}{
			"cardId": params.CardId,
			"name":   params.Name,
			"skill":  params.Skill,
			"effect": params.Effect,
			"reason": params.Reason,
			"remark": params.Remark,
		})
		if result2.Error != nil {
			MyErr(result2.Error.Error(), c)
			return
		}
		result3 := my.DB.Table("password").Where("id = ?", obj.ID).Delete(nil)
		if result3.Error != nil {
			MyErr(result3.Error.Error(), c)
			return
		}
		HandleOk(c, "操作成功")
	}
}

func cardDiyList(c *gin.Context) {
	var data []models.CardDiyAll
	result := my.DB.Table("card_diy").Find(&data)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.CardDiyAll]("查询成功", c, data)
}

func cardDiyAdd(c *gin.Context) {
	var params models.CardDiyBase
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("card_diy").Create(&params)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func cardDiyAddAll(c *gin.Context) {
	var params models.CardDiyAddData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.CardDiyBase
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.Name = item.Name
		data.Zhenyin = item.Zhenyin
		data.Cost = item.Cost
		data.Quality = item.Quality
		data.CardType = item.CardType
		data.Att = item.Att
		data.Life = item.Life
		data.Effect = item.Effect
		data.Info = item.Info
		data.Remark = item.Remark
		data.Img = item.Img
		result := my.DB.Table("card_diy").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}

func cardDiyUpdate(c *gin.Context) {
	var params models.CardDiyUpdate
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	if params.Password != "suola18" {
		MyErr("管理员密码错误", c)
		return
	} else {
		result := my.DB.Table("card_diy").Where("id = ?", params.ID).Updates(map[string]interface{}{
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
		})
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
		HandleOk(c, "操作成功")
	}
}

func cardDiyUpdateTemp(c *gin.Context) {
	var params models.CardDiyUpdate
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var obj models.FrequencyPaddwordAll
	result := my.DB.Table("password").Where("password = ?", params.Password).Find(&obj)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	if result.RowsAffected == 0 {
		MyErr("临时密码错误", c)
		return
	} else {
		result2 := my.DB.Table("card_diy").Where("id = ?", params.ID).Updates(map[string]interface{}{
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
		})
		if result2.Error != nil {
			MyErr(result2.Error.Error(), c)
			return
		}
		result3 := my.DB.Table("password").Where("id = ?", obj.ID).Delete(nil)
		if result3.Error != nil {
			MyErr(result3.Error.Error(), c)
			return
		}
		HandleOk(c, "操作成功")
	}
}
