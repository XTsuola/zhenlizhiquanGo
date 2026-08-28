package models

// BatchData 批量写入请求封装
type BatchData[T any] struct {
	Data []T `json:"data"`
}
