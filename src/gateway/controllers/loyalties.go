package controllers

import (
	"gateway/controllers/responses"
	"gateway/errors"
	"gateway/models"
	"github.com/gorilla/mux"
	"net/http"
)

type loyaltiesController struct {
	loyalties *models.LoyaltyModel
}

func InitLoyalties(r *mux.Router, loyalties *models.LoyaltyModel) {
	controller := &loyaltiesController{loyalties: loyalties}
	r.HandleFunc("/loyalties", controller.getLoyaltyForUser).Methods("GET")
}

func (controller *loyaltiesController) getLoyaltyForUser(w http.ResponseWriter, r *http.Request) {
	loyalty, err := controller.loyalties.GetLoyaltyForUser(r.Header.Get("Authorization"))

	switch err {
	case nil:
		responses.JsonSuccess(w, loyalty)
	case errors.RecordNotFound:
		responses.RecordNotFound(w, r.Header.Get("Authorization"))
	default:
		responses.InternalError(w)
	}
}
