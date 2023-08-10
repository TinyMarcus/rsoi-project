package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type validationErrors struct{}
type validationErrorResponse struct {
	Message string           `json:"message"`
	Errors  validationErrors `json:"errors"`
}

func InternalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode("Internal error")
}

func BadRequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(msg)
}

func ValidationErrorResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)

	response := &validationErrorResponse{"Request validation failed", validationErrors{}}
	json.NewEncoder(w).Encode(response)
}

func RecordNotFound(w http.ResponseWriter, recType string) {
	message := fmt.Sprintf("Not found %s for Id", recType)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(message)
}

func TextSuccess(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(message)
}

func JsonSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")

	json.NewEncoder(w).Encode(data)
}

func SuccessCreation(w http.ResponseWriter, location string) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
}

func SuccessHotelDeletion(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)

	json.NewEncoder(w).Encode("Возврат билета успешно выполнен")
}

func SuccessReservationDeletion(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Response-Code", "04")

	json.NewEncoder(w).Encode(data)
}

func TokenIsMissing(w http.ResponseWriter) {
	message := "Missing auth token"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)

	json.NewEncoder(w).Encode(message)
}

func JwtAccessDenied(w http.ResponseWriter) {
	message := "jwt-token is not valid"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)

	json.NewEncoder(w).Encode(message)
}

func TokenExpired(w http.ResponseWriter) {
	message := "jwt-token expired"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)

	json.NewEncoder(w).Encode(message)
}
