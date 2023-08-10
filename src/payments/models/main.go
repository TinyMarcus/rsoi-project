package models

import (
	"github.com/jinzhu/gorm"
	"payments/repositories"
)

type Models struct {
	Payment *PaymentModel
}

func InitModels(db *gorm.DB) *Models {
	models := new(Models)
	models.Payment = NewPaymentModel(repositories.NewPostgresPaymentRepository(db))

	return models
}
