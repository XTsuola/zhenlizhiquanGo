package models

// PasswordAll 一次性临时密码（表 password）
type PasswordAll struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Password string `json:"password"`
}

func (PasswordAll) TableName() string { return "password" }

// PasswordAdd 新增临时密码
type PasswordAdd struct {
	Password string `json:"password"`
}

func (PasswordAdd) TableName() string { return "password" }

// FrequencyPaddwordAll 兼容旧命名（拼写错误 FrequencyPaddword*）
type FrequencyPaddwordAll = PasswordAll
type FrequencyPaddwordAdd = PasswordAdd
