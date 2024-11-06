package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/suraboy/poc-api-architectures/internal"
	"io/ioutil"
	"net/http"
)

type Service interface {
	GetUser() (UserResponse, error)
}

type service struct {
	client *http.Client
	conf   *internal.Config
}

func NewService(conf *internal.Config) Service {
	return &service{
		client: &http.Client{},
		conf:   conf,
	}
}

func (s service) GetUser() (UserResponse, error) {
	requestBody := []byte(``)
	baseUrl := s.conf.WebClient.RestApi.Endpoint + s.conf.WebClient.RestApi.Paths.UserUrl
	req, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return UserResponse{}, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return UserResponse{}, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return UserResponse{}, fmt.Errorf("failed to read response: %v", err)
	}

	var response UserResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return UserResponse{}, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return response, nil
}
