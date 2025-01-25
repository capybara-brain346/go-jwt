package initializers

import (
	"github.com/capybara-brain346/go-jwt/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
