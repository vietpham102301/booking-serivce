package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type EmployeeResponse struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	DateOfBirth string             `json:"date_of_birth"`
}
