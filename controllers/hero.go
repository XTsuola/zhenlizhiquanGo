package controllers

import (
	my "go_project/config"
	"go_project/models"
	"time"

	"github.com/gin-gonic/gin"
)

func heroList(c *gin.Context) {
	var list []models.HeroSelect
	result := my.DB.Table("hero").Find(&list)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	var data []models.HeroAll
	for _, item := range list {
		var obj models.HeroAll
		obj.ID = item.ID
		obj.Name = item.Name
		obj.Quality = item.Quality
		obj.Zhu = item.Zhu
		obj.Fu = item.Fu
		obj.SkillName = item.SkillName
		obj.Img = item.Img
		obj.Data = StringToArr[models.HeroData](item.Data)
		data = append(data, obj)
	}
	SearchList[models.HeroAll]("查询成功", c, data)
}

func heroAdd(c *gin.Context) {
	var params models.HeroAddData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.HeroAddObj
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.Name = item.Name
		data.Quality = item.Quality
		data.Zhu = item.Zhu
		data.Fu = item.Fu
		data.SkillName = item.SkillName
		data.Img = item.Img
		data.Data = ArrToString[models.HeroData](item.Data)
		result := my.DB.Table("hero").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}

func shardList(c *gin.Context) {
	var list []models.ShardAll
	result := my.DB.Table("shard").Find(&list)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	SearchList[models.ShardAll]("查询成功", c, list)
}

func shardUpdate(c *gin.Context) {
	var params models.ShardUpdateParams
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	result := my.DB.Table("shard").Where("id = ?", params.ID).Updates(map[string]interface{}{
		"skillData": params.SkillData,
	})
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return
	}
	HandleOk(c, "操作成功")
}

func shardAdd(c *gin.Context) {
	var params models.ShardAddData
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return
	}
	var data models.ShardBase
	for _, item := range params.Data {
		time.Sleep(50 * time.Millisecond)
		data.Quality = item.Quality
		data.LevelData = item.LevelData
		data.SkillData = item.SkillData
		result := my.DB.Table("shard").Create(&data)
		if result.Error != nil {
			MyErr(result.Error.Error(), c)
			return
		}
	}
	HandleOk(c, "新增成功")
}
