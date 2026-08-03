package controllers

import (
	"encoding/base64"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	my "go_project/config"
	"go_project/models"
	"go_project/utils"
)

const adminPassword = "suola18"

// ArrToString 数组转字符串（兼容旧调用，实现见 utils）
func ArrToString[T any](arr []T) string {
	return utils.ArrToString(arr)
}

// StringToArr 字符串转数组
func StringToArr[T any](str string) []T {
	return utils.StringToArr[T](str)
}

// WriteImg 写入 base64 图片文件
func WriteImg(imgStr string, savePath string) error {
	if idx := strings.Index(imgStr, ","); idx != -1 {
		imgStr = imgStr[idx+1:]
	}
	data, err := base64.StdEncoding.DecodeString(imgStr)
	if err != nil {
		return err
	}
	return os.WriteFile(savePath, data, 0644)
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func MyErr(err string, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"msg":  err,
	})
}

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

func SearchOne[T any](msg string, c *gin.Context, data T) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  If(msg == "", "success", msg),
		"data": data,
	})
}

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

func HandleOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
	})
}

// bindJSON 绑定请求体，失败时写 MyErr 并返回 false
func bindJSON[T any](c *gin.Context) (T, bool) {
	var params T
	if err := c.ShouldBindJSON(&params); err != nil {
		MyErr(err.Error(), c)
		return params, false
	}
	return params, true
}

func queryInt(c *gin.Context, key string) int {
	v, _ := strconv.Atoi(c.Query(key))
	return v
}

func requireAdminPassword(password string, c *gin.Context) bool {
	if password != adminPassword {
		MyErr("管理员密码错误", c)
		return false
	}
	return true
}

// consumeTempPassword 校验并消费一次性临时密码，返回 password 表 id
func consumeTempPassword(password string, c *gin.Context) (int, bool) {
	var obj models.PasswordAll
	result := my.DB.Table("password").Where("password = ?", password).Find(&obj)
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

func deleteTempPassword(id int, c *gin.Context) bool {
	if err := my.DB.Table("password").Where("id = ?", id).Delete(nil).Error; err != nil {
		MyErr(err.Error(), c)
		return false
	}
	return true
}

// createBatch 批量写入，每条独立 Create，避免复用结构体导致主键污染
func createBatch[T any](table string, rows []T, c *gin.Context) bool {
	for i := range rows {
		if err := my.DB.Table(table).Create(&rows[i]).Error; err != nil {
			MyErr(err.Error(), c)
			return false
		}
	}
	return true
}

func findAll[T any](table string, c *gin.Context) ([]T, bool) {
	var data []T
	if err := my.DB.Table(table).Find(&data).Error; err != nil {
		MyErr(err.Error(), c)
		return nil, false
	}
	return data, true
}

func deleteByQuery(table string, c *gin.Context, query string, args ...interface{}) bool {
	if err := my.DB.Table(table).Where(query, args...).Delete(nil).Error; err != nil {
		MyErr(err.Error(), c)
		return false
	}
	return true
}

func updateByID(table string, id interface{}, values map[string]interface{}, c *gin.Context) bool {
	if err := my.DB.Table(table).Where("id = ?", id).Updates(values).Error; err != nil {
		MyErr(err.Error(), c)
		return false
	}
	return true
}
