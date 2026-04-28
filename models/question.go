package models

type QuestionBase struct {
	Info string `json:"info"`
	Time string `json:"time"`
}

type QuestionAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	QuestionBase
}

type QuestionAddData struct {
	Data []QuestionBase `json:"data"`
}

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

type AnswerAddData struct {
	Data []AnswerBase `json:"data"`
}
