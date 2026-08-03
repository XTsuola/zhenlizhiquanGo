package models

// PageType 分页响应封装
type PageType[T any] struct {
	Rows  []T   `json:"rows"`
	Total int64 `json:"total"`
}

// BatchData 批量写入请求封装
type BatchData[T any] struct {
	Data []T `json:"data"`
}
