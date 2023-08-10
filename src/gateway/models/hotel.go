package models

import (
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

type HotelModel struct {
	client *http.Client
}

func NewHotelModel(client *http.Client) *HotelModel {
	return &HotelModel{
		client: client,
	}
}

func (model *HotelModel) GetHotel(hotelUid string, authHeader string) (*objects.HotelResponseDto, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/hotels/%s", utils.Config.Endpoints.Reservations, hotelUid), nil)
	req.Header.Add("Authorization", authHeader)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode == http.StatusNotFound {
		return nil, errors.RecordNotFound
	} else {
		data := &objects.HotelResponseDto{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, data)
		return data, nil
	}
}

func (model *HotelModel) GetHotels(page int, page_size int, authHeader string) *objects.HotelPaginationResponse {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/hotels", utils.Config.Endpoints.Reservations), nil)
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

	data := &objects.HotelPaginationResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, data)
	return data
}
