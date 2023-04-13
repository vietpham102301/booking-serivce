package booking

import (
	"booking-service/internal/models"
	"booking-service/internal/repos"
	"booking-service/internal/services/sendservice"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestBooking_CreateBooking(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIrepo := repos.NewMockIRepo(mockCtrl)
	mockBookingRepo := repos.NewMockIBookingRepo(mockCtrl)
	mockIrepo.EXPECT().Booking().Return(mockBookingRepo).AnyTimes()

	mockISendService := sendservice.NewMockISendService(mockCtrl)

	var err error
	var emp models.Employee

	mockISendService.EXPECT().GetAvailableEmpl().
		DoAndReturn(func() (*models.EmployeeResponse, error) {
			id, _ := primitive.ObjectIDFromHex("64378dd3181def3832629904")
			return &models.EmployeeResponse{
				ID:          id,
				Name:        "Nguyen Van A",
				DateOfBirth: "01/01/1999",
			}, nil
		}).AnyTimes()

	mockBookingRepo.EXPECT().Create(gomock.Any()).
		DoAndReturn(func(booking *models.Booking) (*models.Booking, error) {
			var availEmp *models.EmployeeResponse
			availEmp, err = mockISendService.GetAvailableEmpl()
			id, _ := primitive.ObjectIDFromHex("64385df47975bfc1dbb0a82b")
			emp = convertEmployee(availEmp)
			expected := &models.Booking{
				ID:          id,
				JobType:     "cleaning the house",
				Description: "come to my house at 8:AM tomorrow",
				Employee:    emp,
			}
			require.Equal(t, expected, booking)

			return booking, nil
		}).AnyTimes()
	b := &Booking{
		mgRepo:         mockIrepo,
		sendServiceAPI: mockISendService,
	}
	oid, _ := primitive.ObjectIDFromHex("64385df47975bfc1dbb0a82b")
	data := &models.Booking{
		ID:          oid,
		JobType:     "cleaning the house",
		Description: "come to my house at 8:AM tomorrow",
	}

	res, err := b.CreateBooking(data)

	expectedRes := &models.Booking{
		ID:          oid,
		JobType:     "cleaning the house",
		Description: "come to my house at 8:AM tomorrow",
		Employee:    emp,
	}
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedRes, res)
}
