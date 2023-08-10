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

type hotelsController struct {
	hotels *models.HotelModel
}

func InitHotels(r *mux.Router, hotels *models.HotelModel) {
	controller := &hotelsController{hotels: hotels}
	r.HandleFunc("/hotels", controller.addHotel).Methods("POST")
	r.HandleFunc("/hotels/{hotelUid}", controller.deleteHotel).Methods("DELETE")
	r.HandleFunc("/hotels", controller.getHotels).Methods("GET")
	r.HandleFunc("/hotels/{hotelUid}", controller.getHotelByUid).Methods("GET")
}

func (controller *hotelsController) addHotel(w http.ResponseWriter, r *http.Request) {
	requestBody := new(objects.CreateHotelRequestDto)
	json.NewDecoder(r.Body).Decode(requestBody)
	reservation, _ := controller.hotels.CreateHotel(requestBody.Name,
		requestBody.Country,
		requestBody.City,
		requestBody.Address,
		requestBody.Stars,
		requestBody.Price)

	responses.JsonSuccess(w, reservation)
}

func (controller *hotelsController) deleteHotel(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	hotelUid := urlParams["hotelUid"]

	err := controller.hotels.DeleteHotel(hotelUid)
	switch err {
	case nil:
		responses.SuccessHotelDeletion(w)
	case errors.RecordNotFound:
		responses.RecordNotFound(w, hotelUid)
	default:
		responses.InternalError(w)
	}
}

func (controller *hotelsController) getHotels(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	pageSize, _ := strconv.Atoi(queryParams.Get("size"))
	hotels := controller.hotels.GetHotels(page, pageSize)

	data := &objects.HotelPaginationResponse{
		Page:          page,
		PageSize:      pageSize,
		TotalElements: len(hotels),
		Items:         objects.ToHotelResponses(hotels),
	}

	responses.JsonSuccess(w, data)
}

func (controller *hotelsController) getHotelByUid(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	hotelUid := urlParams["hotelUid"]
	hotel, err := controller.hotels.GetHotel(hotelUid)

	switch err {
	case nil:
		responses.JsonSuccess(w, hotel.ToHotelResponseDto())
	case errors.RecordNotFound:
		responses.RecordNotFound(w, hotelUid)
	default:
		responses.InternalError(w)
	}
}
