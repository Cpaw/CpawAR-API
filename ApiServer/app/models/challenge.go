package models

type Challenge struct {
	BaseModel
	Title        string `sql:"size:64" json:"title" validate:"min=1, max=64" gorm:"not null"`
	QuestionText string `sql:"size:512" json:"questiontext" validate:"min=1, max=512" gorm:"not null"`
	Weight       float64
}
