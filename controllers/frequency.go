package controllers

import (
	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
	"strconv"
	"time"
)

func frequencyCardsAdd(c *gin.Context) {
	var params models.FrequencyBase
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("frequency").Create(&params)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func frequencyCardsUpdate(c *gin.Context) {
	var params models.FrequencyUpdate
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	if params.Password != "suola18" {
		MyErr("管理员密码错误", c)
		return
	}
	result := my.DB.Table("frequency").Where("id = ?", params.ID).Updates(map[string]interface{}{
		"name":     params.Name,
		"qu":       params.Qu,
		"heroId":   params.HeroId,
		"heroLife": params.HeroLife,
		"cards":    params.Cards,
		"time":     params.Time,
	})
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "操作成功")
}

func frequencyCardsUpdateTemp(c *gin.Context) {
	var params models.FrequencyUpdate
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
		result2 := my.DB.Table("frequency").Where("id = ?", params.ID).Updates(map[string]interface{}{
			"name":     params.Name,
			"qu":       params.Qu,
			"heroId":   params.HeroId,
			"heroLife": params.HeroLife,
			"cards":    params.Cards,
			"time":     params.Time,
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

func frequencyCardsDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var list []models.FrequencyAll
	result := my.DB.Table("frequency").Where("heroId = ?", id).Find(&list)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.FrequencyAll]("查询成功", c, list)
}

func frequencyCardsAll(c *gin.Context) {
	var list []models.FrequencyAll
	result := my.DB.Table("frequency").Find(&list)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.FrequencyAll]("查询成功", c, list)
}

func frequencyCardsDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	password := c.Query("password")
	if password != "suola18" {
		MyErr("管理员密码错误", c)
		return
	}
	result := my.DB.Table("frequency").Where("id = ?", id).Delete(nil)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "删除成功")
}

func frequencyCardsAddAll(c *gin.Context) {
	var params models.FrequencyAddAll
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.FrequencyBase
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.Name = item.Name
		data.Qu = item.Qu
		data.HeroId = item.HeroId
		data.HeroLife = item.HeroLife
		data.Cards = item.Cards
		data.Time = item.Time
		result := my.DB.Table("frequency").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}

func frequencyPasswordAdd(c *gin.Context) {
	var params models.FrequencyPaddwordAdd
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("password").Create(&params)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func frequencyPasswordList(c *gin.Context) {
	var passwordList []models.FrequencyPaddwordAll
	result := my.DB.Table("password").Find(&passwordList)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.FrequencyPaddwordAll]("查询成功", c, passwordList)
}

func frequencyPasswordDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	result := my.DB.Table("password").Where("id = ?", id).Delete(nil)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "删除成功")
}
