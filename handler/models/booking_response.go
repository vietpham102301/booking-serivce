package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookingResponse struct {
	ID          primitive.ObjectID `json:"id"`
	JobType     string             `json:"job_type"`
	Description string             `json:"description"`
	Employee    Employee           `json:"employee"`
}

type Employee struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	DateOfBirth string             `json:"date_of_birth"`
}
