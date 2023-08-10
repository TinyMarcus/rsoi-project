package models

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gateway/errors"
	"gateway/objects"
	"gateway/utils"
	"io/ioutil"
	"net/http"
	"time"
)

type ReservationModel struct {
	client *http.Client
}

func NewReservationModel(client *http.Client) *ReservationModel {
	return &ReservationModel{
		client: client,
	}
}

func (model *ReservationModel) GetReservationsOfUser(authHeader string, page int, page_size int) *objects.ReservationsPaginationResponse {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/reservations", utils.Config.Endpoints.Reservations), nil)
	q := req.URL.Query()
	q.Add("page", fmt.Sprintf("%d", page))
	q.Add("size", fmt.Sprintf("%d", page_size))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", authHeader)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("client: error making http request\n")
	}

	data := &objects.ReservationsPaginationResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, data)
	return data
}

func (model *ReservationModel) GetReservationOfUserByReservationUid(authHeader string, reservationUid string) (*objects.ReservationResponseDto, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/reservations/%s", utils.Config.Endpoints.Reservations, reservationUid), nil)
	req.Header.Add("Authorization", authHeader)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode == http.StatusNotFound {
		return nil, errors.RecordNotFound
	} else {
		data := &objects.ReservationResponseDto{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, data)
		return data, nil
	}
}

// TODO: доделать под логику методы
func (model *ReservationModel) CreateReservation(authHeader string, request *objects.CreateReservationRequestDto) (*objects.ReservationResponseDto, error) {
	requestBody, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/reservations", utils.Config.Endpoints.Reservations), bytes.NewBuffer(requestBody))
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &objects.ReservationResponseDto{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, data)
	return data, nil
}

func (model *ReservationModel) DeleteReservation(authHeader string, reservationUid string) (*objects.ReservationDeletionResponseDto, error) {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/reservations/%s", utils.Config.Endpoints.Reservations, reservationUid), nil)
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &objects.ReservationDeletionResponseDto{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(body, data)
	return data, nil
}
