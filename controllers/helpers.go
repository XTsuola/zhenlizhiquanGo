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

func ArrToString[T any](arr []T) string { return utils.ArrToString(arr) }
func StringToArr[T any](str string) []T { return utils.StringToArr[T](str) }

func adminPassword() string {
	if v := os.Getenv("ADMIN_PASSWORD"); v != "" {
		return v
	}
	return "suola18"
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func MyErr(err string, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": err})
}

func BadRequest(msg string, c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": msg})
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
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": msg})
}

func bindJSON[T any](c *gin.Context) (T, bool) {
	var params T
	if err := c.ShouldBindJSON(&params); err != nil {
		BadRequest(err.Error(), c)
		return params, false
	}
	return params, true
}

func queryInt(c *gin.Context, key string) int {
	v, _ := strconv.Atoi(c.Query(key))
	return v
}

func requireAdminPassword(password string, c *gin.Context) bool {
	if password != adminPassword() {
		MyErr("管理员密码错误", c)
		return false
	}
	return true
}

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

func deleteTempPassword(id int, c *gin.Context) bool {
	if err := my.DB.Table("password").Where("id = ?", id).Delete(&models.PasswordAll{}).Error; err != nil {
		MyErr(err.Error(), c)
		return false
	}
	return true
}

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
	result := my.DB.Table(table).Where("id = ?", id).Updates(values)
	if result.Error != nil {
		MyErr(result.Error.Error(), c)
		return false
	}
	return true
}
