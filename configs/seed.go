// config/seed.go
package config

import (
	"log"
	models "vibex-api/internal/model"
)

func Seed() {

	statuses := []models.Status{
		{Value: "active"},
		{Value: "inactive"},
		{Value: "blocked"},
	}

	for _, status := range statuses {
		var count int64
		DB.Model(&models.Status{}).Where("value = ?", status.Value).Count(&count)

		if count == 0 {

			if err := DB.Create(&status).Error; err != nil {
				log.Fatalf("Failed to seed status %s: %v", status.Value, err)
			}
			log.Printf("Inserted status: %s\n", status.Value)
		} else {
			log.Printf("Status '%s' already exists, skipping insertion.\n", status.Value)
		}
	}
}
