package models

type QuestionBase struct {
	Info string `json:"info"`
	Time string `json:"time"`
}

type QuestionAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	QuestionBase
}

type AnswerBase struct {
	QuestionId int    `json:"questionId" gorm:"column:questionI"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	Time       string `json:"time"`
}
