package repos

import (
	"booking-service/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type BookingNoSQLRepo struct {
	mongoClient *mongo.Client
}

func NewBookingNoSQLRepo(mongoClient *mongo.Client) IBookingRepo {
	return &BookingNoSQLRepo{
		mongoClient: mongoClient,
	}
}

func (b *BookingNoSQLRepo) Create(booking *models.Booking) (*models.Booking, error) {
	bookingCollection := b.mongoClient.Database("booking-service-db").Collection("bookings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := bookingCollection.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}

	return booking, nil
}
