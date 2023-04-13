package repos

import "booking-service/internal/models"

//go:generate mockgen -package=repos -destination=irepo_mock.go -source=irepo.go
type IRepo interface {
	Booking() IBookingRepo
}

type IBookingRepo interface {
	Create(booking *models.Booking) (*models.Booking, error)
}
