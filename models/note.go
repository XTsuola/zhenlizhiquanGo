package models

type NoteBase struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Time    string `json:"time"`
	Content string `json:"content"`
}

type NoteAll struct {
	ID int `json:"id" gorm:"primaryKey"`
	NoteBase
}

// TableName 表名 note
func (NoteAll) TableName() string { return "note" }
