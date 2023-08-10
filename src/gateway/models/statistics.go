package models

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/objects"
	"gateway/utils"
	"io/ioutil"
	"net/http"
	"time"
)

type StatisticsModel struct {
	client *http.Client
}

func NewStatisticsModel(client *http.Client) *StatisticsModel {
	return &StatisticsModel{client: client}
}

func (model *StatisticsModel) Fetch(beginTime time.Time, endTime time.Time, authHeader string) *objects.FetchResponse {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/requests", utils.Config.Endpoints.Statistics), nil)
	q := req.URL.Query()
	q.Add("begin_time", beginTime.Format(time.RFC3339))
	q.Add("end_time", endTime.Format(time.RFC3339))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", authHeader)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("client: error making http request\n")
	}

	data := &objects.FetchResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, data)
	return data
}
