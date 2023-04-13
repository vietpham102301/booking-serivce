package handler

import (
	"booking-service/handler/models"
	models2 "booking-service/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func (h *Handler) createBooking() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("start creating booking")
		input := models.CreateBookingRequest{}
		err := ctx.Bind(&input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		record := &models2.Booking{
			ID:          primitive.NewObjectID(),
			JobType:     input.JobType,
			Description: input.Description,
		}
		record, err = h.booking.CreateBooking(record)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		res := models.BookingResponse{
			ID:          record.ID,
			JobType:     record.JobType,
			Description: record.Description,
			Employee: models.Employee{
				ID:          record.Employee.ID,
				Name:        record.Employee.Name,
				DateOfBirth: record.Employee.DateOfBirth,
			},
		}
		ctx.JSON(http.StatusOK, res)
	}
}
