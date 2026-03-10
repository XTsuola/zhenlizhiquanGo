package controllers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

// WriteImg 写入二进制图片文件封装
func WriteImg(imgStr string, savePath string) {
	if idx := strings.Index(imgStr, ","); idx != -1 {
		imgStr = imgStr[idx+1:]
	}
	data, _ := base64.StdEncoding.DecodeString(imgStr)
	err := os.WriteFile(savePath, data, 0644)
	if err != nil {
		return
	}
}

// If 三元表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

// ArrToString 数组转字符串
func ArrToString[T any](arr []T) string {
	if len(arr) == 0 {
		return `[]`
	} else {
		jsonBytes, _ := json.Marshal(arr)
		jsonStr := string(jsonBytes)
		return jsonStr
	}
}

// StringToArr 字符串转数组
func StringToArr[T any](str string) []T {
	var arr []T
	err := json.Unmarshal([]byte(str), &arr)
	if err != nil || len(arr) == 0 {
		arr = []T{}
	}
	return arr
}

// MyErr 接口500报错
func MyErr(err string, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": 500,
		"msg":  err,
	})
}

// SearchList 查询列表成功
func SearchList[T any](msg string, c *gin.Context, data []T) {
	if len(data) == 0 {
		data = []T{}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  If(msg == "", "success", msg),
		"data": data,
	})
}

// SearchOne 查询单个
func SearchOne[T any](msg string, c *gin.Context, data T) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  If(msg == "", "success", msg),
		"data": data,
	})
}

// SearchByPage 分页查询成功
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

// HandleOk 操作成功
func HandleOk(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
	})
}
