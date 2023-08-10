package models

import (
	"reservations/objects"
	"reservations/repositories"
)

type HotelModel struct {
	repository repositories.HotelRepository
}

func NewHotelModel(repository repositories.HotelRepository) *HotelModel {
	return &HotelModel{
		repository: repository,
	}
}

func (model *HotelModel) GetHotel(hotelUid string) (*objects.Hotel, error) {
	hotel, err := model.repository.Find(hotelUid)
	if err != nil {
		return nil, err
	}

	return hotel, nil
}

func (model *HotelModel) GetHotels(page int, page_size int) []objects.Hotel {
	return model.repository.FindHotels(page, page_size)
}

// TODO: доделать под логику методы
func (model *HotelModel) CreateHotel(name, country, city, address string, stars, price int) (*objects.Hotel, error) {
	hotel := &objects.Hotel{
		Name:    name,
		Country: country,
		City:    city,
		Address: address,
		Stars:   stars,
		Price:   price,
	}
	err := model.repository.Create(hotel)
	return hotel, err
}

func (model *HotelModel) DeleteHotel(hotelUid string) error {
	return model.repository.Delete(hotelUid)
}
