package config

import (
	"log"

	"github.com/mutinho/src/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
	)

	if err != nil {
		log.Fatal("Error migrating database", err)
	}
	log.Println("Database migrated")
}
