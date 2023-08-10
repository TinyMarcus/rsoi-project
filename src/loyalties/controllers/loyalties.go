package controllers

import (
	"github.com/gorilla/mux"
	"loyalties/controllers/responses"
	"loyalties/errors"
	"loyalties/models"
	"loyalties/objects"
	"net/http"
)

type loyaltiesController struct {
	loyalties *models.LoyaltyModel
}

func InitLoyalties(r *mux.Router, loyalties *models.LoyaltyModel) {
	controller := &loyaltiesController{loyalties: loyalties}
	r.HandleFunc("/loyalties", controller.getLoyaltyForUser).Methods("GET")
	r.HandleFunc("/loyalties", controller.increaseLoyalty).Methods("POST")
	r.HandleFunc("/loyalties", controller.decreaseLoyalty).Methods("DELETE")

}

func (controller *loyaltiesController) increaseLoyalty(w http.ResponseWriter, r *http.Request) {
	err := controller.loyalties.IncreaseLoyalty(r.Header.Get("X-User-Name"))
	switch err {
	case nil:
		responses.SuccessCreation(w, "Client's loyalty was created")
	default:
		responses.BadRequest(w, err.Error())
	}
}

func (controller *loyaltiesController) decreaseLoyalty(w http.ResponseWriter, r *http.Request) {
	err := controller.loyalties.DecreaseLoyalty(r.Header.Get("X-User-Name"))
	switch err {
	case nil:
		responses.SuccessCreation(w, "Client's loyalty was created")
	case errors.RecordNotFound:
		responses.RecordNotFound(w, "Record for such user was not found")
	default:
		responses.BadRequest(w, err.Error())
	}
}

func (controller *loyaltiesController) getLoyaltyForUser(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-User-Name")
	loyalty, _ := controller.loyalties.GetLoyaltyForUser(username)

	if loyalty != nil {
		data := objects.ToLoyaltyDtoResponse(loyalty)
		responses.JsonSuccess(w, data)
	} else {
		responses.RecordNotFound(w, "loyalty")
	}
}
