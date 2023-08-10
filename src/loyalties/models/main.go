package models

import (
	"github.com/jinzhu/gorm"
	"loyalties/repositories"
)

type Models struct {
	Loyalty *LoyaltyModel
}

func InitModels(db *gorm.DB) *Models {
	models := new(Models)
	models.Loyalty = NewLoyaltyModel(repositories.NewPostgresLoyaltyRepository(db))

	return models
}
