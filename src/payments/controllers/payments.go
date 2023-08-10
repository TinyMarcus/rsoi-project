package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"payments/controllers/responses"
	"payments/errors"
	"payments/models"
	"payments/objects"
)

type paymentsController struct {
	payments *models.PaymentModel
}

func InitPayments(r *mux.Router, payments *models.PaymentModel) {
	controller := &paymentsController{payments: payments}
	r.HandleFunc("/payments/{paymentUid}", controller.getPayment).Methods("GET")
	r.HandleFunc("/payments", controller.addPayment).Methods("POST")
	r.HandleFunc("/payments/{paymentUid}", controller.deletePayment).Methods("DELETE")
}

func (controller *paymentsController) addPayment(w http.ResponseWriter, r *http.Request) {
	requestBody := new(objects.PaymentDto)
	json.NewDecoder(r.Body).Decode(requestBody)
	payment, _ := controller.payments.CreatePayment(requestBody.Status, requestBody.Price)

	responses.JsonSuccess(w, payment)
}

func (controller *paymentsController) deletePayment(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	paymentUid := urlParams["paymentUid"]

	err := controller.payments.DeletePayment(paymentUid)
	switch err {
	case nil:
		responses.SuccessPaymentDeletion(w)
	case errors.RecordNotFound:
		responses.RecordNotFound(w, paymentUid)
	default:
		responses.InternalError(w)
	}
}

func (controller *paymentsController) getPayment(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	paymentUid := urlParams["paymentUid"]
	payment, err := controller.payments.GetPayment(paymentUid)

	switch err {
	case nil:
		responses.JsonSuccess(w, objects.ToPaymentDtoResponse(payment))
	case errors.RecordNotFound:
		responses.RecordNotFound(w, paymentUid)
	default:
		responses.InternalError(w)
	}
}
