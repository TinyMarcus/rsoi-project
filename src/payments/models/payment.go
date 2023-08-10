package models

import (
	"payments/objects"
	"payments/repositories"
)

type PaymentModel struct {
	repository repositories.PaymentRepository
}

func NewPaymentModel(repository repositories.PaymentRepository) *PaymentModel {
	return &PaymentModel{
		repository: repository,
	}
}

func (model *PaymentModel) GetPayment(paymentUid string) (*objects.Payment, error) {
	payment, err := model.repository.Find(paymentUid)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// TODO: доделать под логику методы
func (model *PaymentModel) CreatePayment(status string, price int) (*objects.Payment, error) {
	payment := &objects.Payment{
		Status: status,
		Price:  price,
	}
	err := model.repository.Create(payment)
	return payment, err
}

func (model *PaymentModel) DeletePayment(paymentUid string) error {
	return model.repository.Delete(paymentUid)
}
