package controllers

import (
	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

func shijiesaiRange(gameType int) (int, int) {
	return gameType * 10000, (gameType + 1) * 10000
}

func shijiesaiList(c *gin.Context) {
	pageSize := queryInt(c, "pageSize")
	pageNo := queryInt(c, "pageNo")
	lo, hi := shijiesaiRange(queryInt(c, "gameType"))
	offset := (pageNo - 1) * pageSize

	var list []models.ShijiesaiAll
	var total int64
	db := my.DB.Table("shijiesai").Where("no > ? AND no < ?", lo, hi)
	if err := db.Count(&total).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	if err := db.Order("no asc").Limit(pageSize).Offset(offset).Find(&list).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}

	data := make([]models.ShijiesaiList, 0, len(list))
	for _, item := range list {
		data = append(data, item.ToList(true))
	}
	SearchByPage("查询成功", c, data, total)
}

func shijiesaiSelect(c *gin.Context) {
	lo, hi := shijiesaiRange(queryInt(c, "gameType"))
	var list []models.ShijiesaiAll
	if err := my.DB.Table("shijiesai").Where("no > ? AND no < ?", lo, hi).Order("no asc").Find(&list).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	data := make([]models.ShijiesaiList, 0, len(list))
	for _, item := range list {
		data = append(data, item.ToList(false))
	}
	SearchList("查询成功", c, data)
}

func shijiesaiAdd(c *gin.Context) {
	params, ok := bindJSON[models.ShijiesaiAdd](c)
	if !ok {
		return
	}
	var exists []models.ShijiesaiAll
	if err := my.DB.Table("shijiesai").Where("no = ?", params.No).Find(&exists).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	if len(exists) != 0 {
		MyErr("该场次已经存在", c)
		return
	}
	data := params.ToBase()
	if err := my.DB.Table("shijiesai").Create(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

func shijiesaiUpdate(c *gin.Context) {
	params, ok := bindJSON[models.ShijiesaiList](c)
	if !ok {
		return
	}
	base := params.ShijiesaiAdd.ToBase()
	if !updateByID("shijiesai", params.ID, map[string]interface{}{
		"no":          base.No,
		"info":        base.Info,
		"shengfuList": base.ShengfuList,
	}, c) {
		return
	}
	HandleOk(c, "操作成功")
}

func shijiesaiDelete(c *gin.Context) {
	if !deleteByQuery("shijiesai", c, "id = ?", queryInt(c, "id")) {
		return
	}
	HandleOk(c, "删除成功")
}

func shijiesaiAddList(c *gin.Context) {
	params, ok := bindJSON[models.ShijiesaiAddData](c)
	if !ok {
		return
	}
	rows := make([]models.ShijiesaiBase, 0, len(params.Data))
	for _, item := range params.Data {
		rows = append(rows, item.ShijiesaiAdd.ToBase())
	}
	if !createBatch("shijiesai", rows, c) {
		return
	}
	HandleOk(c, "新增成功")
}

func shijiesaiAllList(c *gin.Context) {
	var list []models.ShijiesaiAll
	var total int64
	db := my.DB.Table("shijiesai")
	if err := db.Count(&total).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	if err := db.Order("no asc").Find(&list).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}

	data := make([]models.ShijiesaiList, 0, len(list))
	for _, item := range list {
		data = append(data, item.ToList(true))
	}
	SearchByPage("查询成功", c, data, total)
}
