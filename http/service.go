package http

import (
	"encoding/json"
	"io"
	"net/http"
)

type HttpService struct {
	client http.Client
}

func NewHttpService(client http.Client) HttpService {
	return HttpService{client: client}
}

func (h HttpService) MakeRequest(method, url string, response interface{}) error {
	req, err := createHttpRequest(method, url)
	if err != nil {
		return err
	}
	res, err := h.client.Do(req)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}
	return err
}

func createHttpRequest(method, url string) (*http.Request, error) {
	return http.NewRequest(method, url, nil)
}
