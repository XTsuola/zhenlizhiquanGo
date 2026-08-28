package controllers

import (
	"go_project/models"
	"go_project/qingshu"

	"github.com/gin-gonic/gin"
)

func qingshuGetMap(c *gin.Context) {
	mapObj, err := qingshu.LoadOrInit()
	if err != nil {
		MyErr(err.Error(), c)
		return
	}
	SearchOne("查询成功", c, mapObj.ToData())
}

func qingshuReset(c *gin.Context) {
	if err := qingshu.Reset(0); err != nil {
		MyErr(err.Error(), c)
		return
	}
	HandleOk(c, "重置成功")
}

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
