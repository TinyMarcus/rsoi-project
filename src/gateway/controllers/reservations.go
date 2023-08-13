package controllers

import (
	"encoding/json"
	"gateway/controllers/responses"
	"gateway/errors"
	"gateway/models"
	"gateway/objects"
	"strconv"
	"time"

	"net/http"

	"github.com/gorilla/mux"
)

// TODO: НУЖНО СДЕЛАТЬ ЛОГИКУ ПОД ОПЛАТУ И РЕЗЕРВАЦИЮ КАК НА C#
type reservationsController struct {
	hotels       *models.HotelModel
	loyalties    *models.LoyaltyModel
	reservations *models.ReservationModel
	payments     *models.PaymentModel
}

func InitReservations(r *mux.Router, reservations *models.ReservationModel, payments *models.PaymentModel,
	loyalties *models.LoyaltyModel, hotels *models.HotelModel) {
	controller := &reservationsController{
		reservations: reservations,
		payments:     payments,
		loyalties:    loyalties,
		hotels:       hotels,
	}
	r.HandleFunc("/me", controller.me).Methods("GET")
	r.HandleFunc("/reservations", controller.bookHotel).Methods("POST")
	r.HandleFunc("/reservations/{reservationUid}", controller.cancelHotel).Methods("DELETE")
	r.HandleFunc("/reservations", controller.getReservationsOfUser).Methods("GET")
	r.HandleFunc("/reservations/{reservationUid}", controller.getReservationOfUserByUid).Methods("GET")
}

func (controller *reservationsController) bookHotel(w http.ResponseWriter, r *http.Request) {
	requestBody := new(objects.CreateReservationRequestDto)
	err := json.NewDecoder(r.Body).Decode(requestBody)
	if err != nil {
		responses.ValidationErrorResponse(w, err.Error())
		return
	}

	// Check if hotel exists
	hotel, err := controller.hotels.GetHotel(requestBody.HotelUid, r.Header.Get("Authorization"))
	if err != nil {
		responses.RecordNotFound(w, err.Error())
		return
	}

	// Count quantity of nights and price
	startTime, err := time.Parse(time.RFC3339, requestBody.StartDate)
	if err != nil {
		responses.ValidationErrorResponse(w, err.Error())
		return
	}

	endTime, err := time.Parse(time.RFC3339, requestBody.EndDate)
	if err != nil {
		responses.ValidationErrorResponse(w, err.Error())
		return
	}

	timeInterval := endTime.Sub(startTime)
	totalPrice := hotel.Price * int(timeInterval.Hours()/24)
	if totalPrice < 0 {
		responses.ValidationErrorResponse(w, errors.InvalidRequest.Error())
		return
	}

	// Get discount value and status
	loyalty, err := controller.loyalties.GetLoyaltyForUser(r.Header.Get("Authorization"))
	if err != nil {
		responses.RecordNotFound(w, err.Error())
		return
	}

	// Count final price
	totalPriceWithDiscount := totalPrice * (100 - loyalty.Discount) / 100

	// Create payment new record
	newPayment := &objects.PaymentDto{
		Status: "PAID",
		Price:  totalPriceWithDiscount,
	}

	paymentDtoWithUid, err := controller.payments.CreatePayment(r.Header.Get("Authorization"), newPayment)
	if err != nil {
		responses.ValidationErrorResponse(w, err.Error())
		return
	}

	// Increment count in loyalties
	err = controller.loyalties.IncreaseLoyalty(r.Header.Get("Authorization"))
	if err != nil {
		responses.ValidationErrorResponse(w, err.Error())
		return
	}

	// Create new record of reservation
	newReservation := &objects.CreateReservationRequestDto{
		PaymentUid: paymentDtoWithUid.PaymentUid,
		HotelUid:   requestBody.HotelUid,
		Status:     "PAID",
		StartDate:  requestBody.StartDate,
		EndDate:    requestBody.EndDate,
	}
	createdReservation, err := controller.reservations.CreateReservation(r.Header.Get("Authorization"), newReservation)
	if err != nil {
		responses.ValidationErrorResponse(w, err.Error())
		return
	}
	hotelInfoDto := &objects.HotelInfoDto{
		HotelUid:    hotel.HotelUid,
		Name:        hotel.Name,
		FullAddress: hotel.Country + ", " + hotel.City + ", " + hotel.Address,
		Stars:       hotel.Stars,
	}

	(*createdReservation).Hotel = *hotelInfoDto
	(*createdReservation).Payment = *paymentDtoWithUid

	responses.JsonSuccess(w, createdReservation)
}

func (controller *reservationsController) cancelHotel(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	reservationUid := urlParams["reservationUid"]

	// Check if reservation exists
	reservation, err := controller.reservations.GetReservationOfUserByReservationUid(r.Header.Get("Authorization"), reservationUid)
	if err != nil {
		responses.RecordNotFound(w, err.Error())
		return
	}
	if reservation.Status == "CANCELED" {
		responses.BadRequest(w, "This reservation have been already canceled")
		return
	}

	// Set reservation as CANCELLED
	result, err := controller.reservations.DeleteReservation(r.Header.Get("Authorization"), reservationUid)
	if err != nil {
		responses.RecordNotFound(w, err.Error())
		return
	}

	// Set payment as CANCELLED
	err = controller.payments.DeletePayment(r.Header.Get("Authorization"), result.PaymentUid)
	if err != nil {
		responses.RecordNotFound(w, err.Error())
		return
	}

	// Decrement count in loyalties
	err = controller.loyalties.DecreaseLoyalty(r.Header.Get("Authorization"))
	if err != nil {
		responses.ValidationErrorResponse(w, err.Error())
		return
	}

	responses.SuccessReservationDeletion(w)
}

func (controller *reservationsController) getReservationsOfUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	pageSize, _ := strconv.Atoi(queryParams.Get("size"))
	if page == 0 && pageSize == 0 {
		page, pageSize = 1, 9999
	}
	reservations := controller.reservations.GetReservationsOfUser(r.Header.Get("Authorization"), page, pageSize)

	for idx, val := range reservations.Items {
		paymentDtoWithId, _ := controller.payments.GetPayment(r.Header.Get("Authorization"), val.Payment.PaymentUid)
		(*reservations).Items[idx].Payment = *paymentDtoWithId
	}

	responses.JsonSuccess(w, reservations.Items)
}

func (controller *reservationsController) getReservationOfUserByUid(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	reservationUid := urlParams["reservationUid"]

	reservation, err := controller.reservations.GetReservationOfUserByReservationUid(r.Header.Get("Authorization"), reservationUid)
	if err != nil {
		responses.RecordNotFound(w, err.Error())
		return
	}

	paymentDtoWithId, _ := controller.payments.GetPayment(r.Header.Get("Authorization"), reservation.Payment.PaymentUid)
	(*reservation).Payment = *paymentDtoWithId

	responses.JsonSuccess(w, reservation)
}

func (controller *reservationsController) me(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	pageSize, _ := strconv.Atoi(queryParams.Get("size"))
	if page == 0 && pageSize == 0 {
		page, pageSize = 1, 9999
	}

	loyalty, err := controller.loyalties.GetLoyaltyForUser(r.Header.Get("Authorization"))
	if err != nil {
		responses.ValidationErrorResponse(w, err.Error())
		return
	}

	reservations := controller.reservations.GetReservationsOfUser(r.Header.Get("Authorization"), page, pageSize)

	userInfo := &objects.UserInfoResponse{
		Loyalties:    *loyalty,
		Reservations: *reservations,
	}

	responses.JsonSuccess(w, userInfo)
}
