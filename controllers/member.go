package controllers

import (
	my "go_project/config"
	"go_project/models"

	"github.com/gin-gonic/gin"
)

// memberList 成员列表，按捐赠和积分降序
func memberList(c *gin.Context) {
	var data []models.MemberAll
	if err := my.DB.Table("member").Order("donation desc, score desc").Find(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchList("查询成功", c, data)
}

// memberAdd 新增成员
func memberAdd(c *gin.Context) {
	params, ok := bindJSON[models.MemberBase](c)
	if !ok {
		return
	}
	if err := my.DB.Table("member").Create(&params).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "新增成功")
}

// memberUpdate 按路径中的 ID 修改成员
func memberUpdate(c *gin.Context) {
	params, ok := bindJSON[models.MemberBase](c)
	if !ok {
		return
	}
	if !updateByID("member", c.Param("id"), map[string]interface{}{
		"name":     params.Name,
		"donation": params.Donation,
		"reward":   params.Reward,
		"score":    params.Score,
		"title":    params.Title,
		"remark":   params.Remark,
	}, c) {
		return
	}
	HandleOk(c, "操作成功")
}

// memberDelete 按路径中的 ID 删除成员
func memberDelete(c *gin.Context) {
	if !deleteByQuery("member", c, "id = ?", c.Param("id")) {
		return
	}
	HandleOk(c, "删除成功")
}

// memberRewardList 奖励大于 0 的成员，按奖励降序
func memberRewardList(c *gin.Context) {
	var data []models.MemberAll
	if err := my.DB.Table("member").Where("reward > ?", 0).Order("reward desc").Find(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchList("查询成功", c, data)
}

// memberAddAll 批量新增成员
func memberAddAll(c *gin.Context) {
	params, ok := bindJSON[models.BatchData[models.MemberBase]](c)
	if !ok {
		return
	}
	if !createBatch("member", params.Data, c) {
		return
	}
	HandleOk(c, "新增成功")
}
