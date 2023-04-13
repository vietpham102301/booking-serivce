package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Booking struct {
	ID          primitive.ObjectID `bson:"_id"`
	JobType     string             `bson:"job_type"`
	Description string             `bson:"description"`
	Employee    Employee           `bson:"employee"`
}

type Employee struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	DateOfBirth string             `bson:"date_of_birth"`
}
