package utils

import "encoding/json"

// ArrToString 数组转 JSON 字符串（空数组返回 "[]"）
func ArrToString[T any](arr []T) string {
	if len(arr) == 0 {
		return `[]`
	}
	jsonBytes, err := json.Marshal(arr)
	if err != nil {
		return `[]`
	}
	return string(jsonBytes)
}

// StringToArr JSON 字符串转数组（解析失败返回空切片）
func StringToArr[T any](str string) []T {
	var arr []T
	if str == "" {
		return arr
	}
	if err := json.Unmarshal([]byte(str), &arr); err != nil || len(arr) == 0 {
		return []T{}
	}
	return arr
}
