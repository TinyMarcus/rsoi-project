package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"reservations/controllers/responses"
	"reservations/errors"
	"reservations/models"
	"reservations/objects"
	"strconv"
)

type reservationsController struct {
	reservations *models.ReservationModel
	hotels       *models.HotelModel
}

func InitReservations(r *mux.Router, reservations *models.ReservationModel, hotels *models.HotelModel) {
	controller := &reservationsController{reservations: reservations, hotels: hotels}
	r.HandleFunc("/reservations", controller.addReservation).Methods("POST")
	r.HandleFunc("/reservations/{reservationUid}", controller.deleteReservation).Methods("DELETE")
	r.HandleFunc("/reservations", controller.getReservationsOfUser).Methods("GET")
	r.HandleFunc("/reservations/{reservationUid}", controller.getReservationOfUserByUid).Methods("GET")
}

func (controller *reservationsController) addReservation(w http.ResponseWriter, r *http.Request) {
	requestBody := new(objects.CreateReservationRequestDto)
	username := r.Header.Get("X-User-Name")
	json.NewDecoder(r.Body).Decode(requestBody)
	reservation, _ := controller.reservations.CreateReservation(username,
		requestBody.PaymentUid,
		requestBody.HotelUid,
		requestBody.Status,
		requestBody.StartDate,
		requestBody.EndDate)

	responses.JsonSuccess(w, reservation)
}

func (controller *reservationsController) deleteReservation(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	reservationUid := urlParams["reservationUid"]
	username := r.Header.Get("X-User-Name")

	paymentUid, err := controller.reservations.DeleteReservation(username, reservationUid)
	responseData := &objects.ReservationDeletionResponseDto{PaymentUid: paymentUid}
	switch err {
	case nil:
		responses.SuccessReservationDeletion(w, responseData)
	case errors.RecordNotFound:
		responses.RecordNotFound(w, reservationUid)
	default:
		responses.InternalError(w)
	}
}

func (controller *reservationsController) getReservationsOfUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	pageSize, _ := strconv.Atoi(queryParams.Get("size"))
	username := r.Header.Get("X-User-Name")
	reservations := controller.reservations.GetReservationsOfUser(username, page, pageSize)

	hotels := make([]objects.HotelInfoDto, len(reservations))
	payments := make([]objects.PaymentDtoWithId, len(reservations))
	for k, v := range reservations {
		hotel, _ := controller.hotels.GetHotel(v.HotelUid)
		hotels[k] = *objects.ToHotelInfoDto(hotel)
		payments[k] = objects.PaymentDtoWithId{PaymentUid: v.PaymentUid}
	}

	data := &objects.ReservationsPaginationResponse{
		Page:          page,
		PageSize:      pageSize,
		TotalElements: len(hotels),
		Items:         objects.ToReservationResponses(reservations, hotels, payments),
	}

	responses.JsonSuccess(w, data)
}

func (controller *reservationsController) getReservationOfUserByUid(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	reservationUid := urlParams["reservationUid"]
	username := r.Header.Get("X-User-Name")
	reservation, err := controller.reservations.GetReservationOfUserByReservationUid(username, reservationUid)
	hotel, err := controller.hotels.GetHotel(reservation.HotelUid)
	hotelDto := objects.ToHotelInfoDto(hotel)
	paymentDtoWithId := &objects.PaymentDtoWithId{PaymentUid: reservation.PaymentUid}

	switch err {
	case nil:
		responses.JsonSuccess(w, objects.ToReservationResponseDto(reservation, hotelDto, paymentDtoWithId))
	case errors.RecordNotFound:
		responses.RecordNotFound(w, reservationUid)
	default:
		responses.InternalError(w)
	}
}
