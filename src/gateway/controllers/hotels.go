package controllers

import (
	"gateway/controllers/responses"
	"gateway/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type hotelsController struct {
	hotels *models.HotelModel
}

func InitHotels(r *mux.Router, hotels *models.HotelModel) {
	controller := &hotelsController{hotels: hotels}
	r.HandleFunc("/hotels", controller.getHotels).Methods("GET")
}

func (controller *hotelsController) getHotels(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	page, _ := strconv.Atoi(queryParams.Get("page"))
	pageSize, _ := strconv.Atoi(queryParams.Get("size"))
	if page == 0 && pageSize == 0 {
		page, pageSize = 1, 9999
	}
	hotels := controller.hotels.GetHotels(page, pageSize, r.Header.Get("Authorization"))

	responses.JsonSuccess(w, hotels)
}
