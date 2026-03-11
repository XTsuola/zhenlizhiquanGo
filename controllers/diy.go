package controllers

import (
	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
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
