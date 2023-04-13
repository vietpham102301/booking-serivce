package booking

import (
	"booking-service/internal/models"
	"booking-service/internal/repos"
	"booking-service/internal/services/sendservice"
	"log"
)

type IBooking interface {
	CreateBooking(booking *models.Booking) (*models.Booking, error)
}

type Booking struct {
	mgRepo         repos.IRepo
	sendServiceAPI sendservice.ISendService
}

func NewBooking(repo repos.IRepo, sendServiceAPI sendservice.ISendService) IBooking {
	return &Booking{
		mgRepo:         repo,
		sendServiceAPI: sendServiceAPI,
	}
}

func (b Booking) CreateBooking(booking *models.Booking) (*models.Booking, error) {
	availEmployee, err := b.sendServiceAPI.GetAvailableEmpl()
	if err != nil {
		log.Printf("create booking fail with err %v", err)
		return nil, err
	}
	emp := convertEmployee(availEmployee)

	booking.Employee = emp

	record, err := b.mgRepo.Booking().Create(booking)
	if err != nil {
		log.Printf("create booking fail with err %v", err)
		return nil, err
	}
	return record, nil
}
