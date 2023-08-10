package models

import (
	"github.com/jinzhu/gorm"
	"reservations/repositories"
)

type Models struct {
	Reservation *ReservationModel
	Hotel       *HotelModel
}

func InitModels(db *gorm.DB) *Models {
	models := new(Models)
	models.Reservation = NewReservationModel(repositories.NewPostgresReservationRepository(db))
	models.Hotel = NewHotelModel(repositories.NewPostgresHotelRepository(db))

	return models
}
