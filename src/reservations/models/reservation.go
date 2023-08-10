package models

import (
	"reservations/objects"
	"reservations/repositories"
)

type ReservationModel struct {
	reservationRepository repositories.ReservationRepository
}

func NewReservationModel(repository repositories.ReservationRepository) *ReservationModel {
	return &ReservationModel{
		reservationRepository: repository,
	}
}

func (model *ReservationModel) GetReservationsOfUser(username string, page int, page_size int) []objects.Reservation {
	reservation := model.reservationRepository.Find(username, page, page_size)

	return reservation
}

func (model *ReservationModel) GetReservationOfUserByReservationUid(username, reservationUid string) (*objects.Reservation, error) {
	reservation, err := model.reservationRepository.FindByReservationUid(username, reservationUid)
	if err != nil {
		return nil, err
	}

	return reservation, nil
}

// TODO: доделать под логику методы
func (model *ReservationModel) CreateReservation(username string,
	paymentUid string,
	hotelUid string,
	reservationStatus string,
	startDate string,
	endDate string) (*objects.Reservation, error) {
	reservation := &objects.Reservation{
		Username:   username,
		PaymentUid: paymentUid,
		HotelUid:   hotelUid,
		Status:     reservationStatus,
		StartDate:  startDate,
		EndDate:    endDate,
	}
	err := model.reservationRepository.Create(reservation)
	return reservation, err
}

func (model *ReservationModel) DeleteReservation(username, reservationUid string) (string, error) {
	reservation, _ := model.reservationRepository.FindByReservationUid(username, reservationUid)

	return reservation.PaymentUid, model.reservationRepository.Delete(reservationUid)
}
