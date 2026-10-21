package models

type QuestionBase struct {
	Info string `json:"info"`
	Time string `json:"time"`
}

type QuestionAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	QuestionBase
}

// TableName 表名 question
func (QuestionAll) TableName() string { return "question" }

type AnswerBase struct {
	QuestionId int    `json:"questionId" gorm:"column:questionId"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	Time       string `json:"time"`
}

type AnswerAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	AnswerBase
}

// TableName 表名 answer
func (AnswerAll) TableName() string { return "answer" }
