package sendservice

import "booking-service/internal/models"

//go:generate mockgen -package=sendservice -destination=sendservice_mock.go -source=isend_service.go

type ISendService interface {
	GetAvailableEmpl() (*models.EmployeeResponse, error)
}
