package handler

import (
	"booking-service/internal/services/booking"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	booking booking.IBooking
}

func NewHandler(booking booking.IBooking) *Handler {
	return &Handler{
		booking: booking,
	}
}

func (h *Handler) ConfigAPIRoute(router *gin.Engine) {
	routers := router.Group("v1")
	routers.POST("booking", h.createBooking())
}
