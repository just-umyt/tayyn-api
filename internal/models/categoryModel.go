package models

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
	Blug []Blug
}
