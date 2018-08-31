package models

type Notification struct {
	BaseModel
	Title        string `sql:"size:64" json:"title" validate:"min=1, max=64" gorm:"not null"`
	Notification string `sql:"size:512" json:"notification" validate:"min=1, max=2048" gorm:"not null"`
}
