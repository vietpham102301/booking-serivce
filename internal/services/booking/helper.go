package booking

import "booking-service/internal/models"

func convertEmployee(response *models.EmployeeResponse) models.Employee {
	return models.Employee{
		ID:          response.ID,
		Name:        response.Name,
		DateOfBirth: response.DateOfBirth,
	}
}
