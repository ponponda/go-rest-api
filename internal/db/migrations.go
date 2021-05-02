package db

import (
	"to-do-api/internal/db/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})
}
