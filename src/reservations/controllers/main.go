package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
	"net/http"
	"reservations/models"
	"reservations/utils"
)

func initControllers(r *mux.Router, m *models.Models) {
	r.Use(utils.LogHandler)
	api1_r := r.PathPrefix("/api/v1/").Subrouter()
	api1_r_noauth := api1_r.NewRoute().Subrouter()
	InitHotels(api1_r_noauth, m.Hotel)
	api1_r_auth := api1_r.NewRoute().Subrouter()
	api1_r_auth.Use(JwtAuthentication)

	InitReservations(api1_r_auth, m.Reservation, m.Hotel)
}

func InitRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	models := models.InitModels(db)

	initControllers(router, models)
	return router
}

func RunRouter(r *mux.Router, port uint16) error {
	c := cors.New(cors.Options{})
	handler := c.Handler(r)
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), handler)
}
