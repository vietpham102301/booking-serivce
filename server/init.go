package server

import (
	handler "booking-service/handler"
	"booking-service/internal/repos"
	"booking-service/internal/services/booking"
	"booking-service/internal/services/sendservice"
)

func (s *Server) initServices(repo repos.IRepo) *ServiceList {

	sendServiceAPI := sendservice.NewSendService()
	booking := booking.NewBooking(repo, sendServiceAPI)

	return &ServiceList{
		booking: booking,
	}
}

func (s *Server) initRouters(serviceList *ServiceList) {
	handler := handler.NewHandler(serviceList.booking)

	handler.ConfigAPIRoute(s.router)
}
