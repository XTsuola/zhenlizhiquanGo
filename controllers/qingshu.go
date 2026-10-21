package controllers

import (
	"go_project/models"
	"go_project/qingshu"

	"github.com/gin-gonic/gin"
)

// qingshuGetMap 获取情书对局，不存在时自动建局
func qingshuGetMap(c *gin.Context) {
	mapObj, err := qingshu.LoadOrInit()
	if err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchOne("查询成功", c, mapObj.ToData())
}

// qingshuReset 重置情书对局
func qingshuReset(c *gin.Context) {
	if err := qingshu.Reset(0); err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "重置成功")
}

// userNameUpdate 修改情书玩家昵称
func userNameUpdate(c *gin.Context) {
	params, ok := bindJSON[models.UsernameUpdate](c)
	if !ok {
		return
	}
	if err := qingshu.UpdateUsername(params.Name, params.Password); err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "操作成功")
}
