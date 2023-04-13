package models

type CreateBookingRequest struct {
	JobType     string `json:"job_type"`
	Description string `json:"description"`
}
