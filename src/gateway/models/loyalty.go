package models

import (
	"encoding/json"
	"fmt"
	"gateway/errors"
	"gateway/objects"
	"gateway/utils"
	"io/ioutil"
	"net/http"
)

type LoyaltyModel struct {
	client *http.Client
}

func NewLoyaltyModel(client *http.Client) *LoyaltyModel {
	return &LoyaltyModel{
		client: client,
	}
}

func (model *LoyaltyModel) GetLoyaltyForUser(authHeader string) (*objects.LoyaltyDto, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/loyalties", utils.Config.Endpoints.Loyalties), nil)
	req.Header.Add("Authorization", authHeader)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode == http.StatusNotFound {
		return nil, errors.RecordNotFound
	} else {
		data := &objects.LoyaltyDto{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, data)
		return data, nil
	}
}

func (model *LoyaltyModel) IncreaseLoyalty(authHeader string) error {
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/loyalties", utils.Config.Endpoints.Loyalties), nil)
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	_, err := client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (model *LoyaltyModel) DecreaseLoyalty(authHeader string) error {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/loyalties", utils.Config.Endpoints.Loyalties), nil)
	req.Header.Add("Authorization", authHeader)
	client := &http.Client{}
	_, err := client.Do(req)
	return err
}
