package models

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(255)"`
	Email      string `gorm:"unique;type:varchar(255)"`
	Password   string
	Nick       string `gorm:"unique;type:varchar(255)"`
	About      string
	Blug       []Blug
	LikedBlugs []Blug `gorm:"many2many:likes;"`
}
