package models

type Blug struct {
	ID         uint `gorm:"primaryKey"`
	Title      string
	Content    string
	UserId     uint
	CategoryId uint
	Likes      []User `gorm:"many2many:likes;"`
}
