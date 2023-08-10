package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/errors"
	"gateway/objects"
	"gateway/utils"
	"io/ioutil"
	"net/http"
)

type PaymentModel struct {
	client *http.Client
}

func NewPaymentModel(client *http.Client) *PaymentModel {
	return &PaymentModel{
		client: client,
	}
}

func (model *PaymentModel) GetPayment(authHeader string, paymentUid string) (*objects.PaymentDtoWithId, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/payments/%s", utils.Config.Endpoints.Payments, paymentUid), nil)
	req.Header.Add("Authorization", authHeader)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode == http.StatusNotFound {
		return nil, errors.RecordNotFound
	} else {
		data := &objects.PaymentDtoWithId{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, data)
		return data, nil
	}
}

// TODO: доделать под логику методы
func (model *PaymentModel) CreatePayment(authHeader string, request *objects.PaymentDto) (*objects.PaymentDtoWithId, error) {
	requestBody, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/payments", utils.Config.Endpoints.Payments), bytes.NewBuffer(requestBody))
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &objects.PaymentDtoWithId{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, data)
	return data, nil
}

func (model *PaymentModel) DeletePayment(authHeader string, paymentUid string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/payments/%s", utils.Config.Endpoints.Payments, paymentUid), nil)
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	_, err := client.Do(req)
	return err
}
