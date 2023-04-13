package sendservice

import (
	"booking-service/internal/models"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

type SendService struct {
	httpClient *http.Client
}

func NewSendService() ISendService {
	return &SendService{
		httpClient: &http.Client{
			Timeout: 120 * time.Second,
		},
	}
}

func (s SendService) GetAvailableEmpl() (*models.EmployeeResponse, error) {
	log.Printf("get employee...")
	url := "http://localhost:3001/v1/employee"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Get employee io.ReadAll error: %v", err)
		return nil, err
	}
	res := &models.EmployeeResponse{}
	if err := json.Unmarshal(body, res); err != nil {
		log.Printf("GET employee json.Unmarshal error: %v", err)
		return nil, err
	}

	return res, nil
}
