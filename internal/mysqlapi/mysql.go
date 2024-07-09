package mysqlapi

import (
	"log"
	"os"

	"github.com/just-umyt/blUg/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB: ", err)
	}

	db.AutoMigrate(&models.Category{}, &models.Blug{}, &models.User{})

	return db
}
