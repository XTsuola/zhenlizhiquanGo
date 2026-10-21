package controllers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
	"go_project/utils"
	"gorm.io/gorm"
)

// ArrToString 将切片序列化为 JSON 字符串；空切片或序列化失败时为 "[]"
func ArrToString[T any](arr []T) string { return utils.ArrToString(arr) }

// StringToArr 将 JSON 字符串反序列化为切片；空串或解析失败时返回空切片
func StringToArr[T any](str string) []T { return utils.StringToArr[T](str) }

// adminPassword 读取管理员密码，优先环境变量 ADMIN_PASSWORD，否则用默认值
func adminPassword() string {
	if v := os.Getenv("ADMIN_PASSWORD"); v != "" {
		return v
	}
	return "suola18"
}

// If 按条件在两个值中择一返回
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// MyErr 返回 HTTP 500，业务码 500
func MyErr(err string, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": err})
}

// BadRequest 返回 HTTP 400，业务码 400
func BadRequest(msg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": msg})
}

// SearchList 返回列表成功响应；data 为 nil 时改为空切片，msg 为空时用 "success"
func SearchList[T any](msg string, c *gin.Context, data []T) {
	if data == nil {
		data = []T{}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  If(msg == "", "success", msg),
		"data": data,
	})
}

// SearchOne 返回单条成功响应；msg 为空时用 "success"
func SearchOne[T any](msg string, c *gin.Context, data T) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  If(msg == "", "success", msg),
		"data": data,
	})
}

// SearchByPage 返回分页成功响应，附带 total；data 为 nil 时改为空切片
func SearchByPage[T any](msg string, c *gin.Context, data []T, total int64) {
	if data == nil {
		data = []T{}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   If(msg == "", "success", msg),
		"data":  data,
		"total": total,
	})
}

// HandleOk 返回无 data 的成功响应
func HandleOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": msg})
}

// bindJSON 绑定 JSON 请求体；失败时已写入 400 响应，第二个返回值为 false
func bindJSON[T any](c *gin.Context) (T, bool) {
	var params T
	if err := c.ShouldBindJSON(&params); err != nil {
		BadRequest(err.Error(), c)
		return params, false
	}
	return params, true
}

// queryInt 读取查询参数并转为 int；缺失或非法时为 0
func queryInt(c *gin.Context, key string) int {
	v, _ := strconv.Atoi(c.Query(key))
	return v
}

// requireAdminPassword 校验管理员密码；不匹配时写入错误响应并返回 false
func requireAdminPassword(password string, c *gin.Context) bool {
	if password != adminPassword() {
		MyErr("管理员密码错误", c)
		return false
	}
	return true
}

// consumeTempPassword 按密码查找临时密码记录并返回其 ID；不存在或查询失败时已写错误响应
func consumeTempPassword(password string, c *gin.Context) (int, bool) {
	var obj models.PasswordAll
	result := my.DB.Table("password").Where("password = ?", password).Limit(1).Find(&obj)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return 0, false
	}
	if result.RowsAffected == 0 {
		MyErr("临时密码错误", c)
		return 0, false
	}
	return obj.ID, true
}

// deleteTempPassword 按 ID 删除已使用的临时密码
func deleteTempPassword(id int, c *gin.Context) bool {
	if err := my.DB.Table("password").Where("id = ?", id).Delete(&models.PasswordAll{}).Error; err != nil {
		MyErr(err.Error(), c)
		return false
	}
	return true
}

// createBatch 在事务中逐条插入；空切片直接成功，任一条失败则整批回滚并写错误响应
func createBatch[T any](table string, rows []T, c *gin.Context) bool {
	if len(rows) == 0 {
		return true
	}
	err := my.DB.Transaction(func(tx *gorm.DB) error {
		for i := range rows {
			if err := tx.Table(table).Create(&rows[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		MyErr(err.Error(), c)
		return false
	}
	return true
}

// findAll 查询表内全部记录；查询失败时已写错误响应
func findAll[T any](table string, c *gin.Context) ([]T, bool) {
	var data []T
	if err := my.DB.Table(table).Find(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return nil, false
	}
	return data, true
}

// deleteByQuery 按条件删除记录；失败时已写错误响应
func deleteByQuery(table string, c *gin.Context, query string, args ...interface{}) bool {
	if err := my.DB.Table(table).Where(query, args...).Delete(nil).Error; err != nil {
		MyErr(err.Error(), c)
		return false
	}
	return true
}

// updateByID 按主键更新指定字段；失败时已写错误响应
func updateByID(table string, id interface{}, values map[string]interface{}, c *gin.Context) bool {
	result := my.DB.Table(table).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return false
	}
	return true
}
