package usecases

import (
	"github.com/disebud/reservation-hotel/main/master/models"
	"github.com/disebud/reservation-hotel/main/master/repositories"
	"github.com/disebud/reservation-hotel/main/master/utils"
)

type BookingUsecaseImpl struct {
	BookingRepo repositories.BookingRepository
}

func InitBookingUsecase(BookingRepo repositories.BookingRepository) BookingUseCase {
	return &BookingUsecaseImpl{BookingRepo}
}

func (s BookingUsecaseImpl) GetAllBookings() ([]*models.Booking, error) {
	Booking, err := s.BookingRepo.GetAllBooking()
	if err != nil {
		return nil, err
	}

	return Booking, nil
}

func (s BookingUsecaseImpl) CreateBooking(Booking models.Booking) error {
	err := utils.ValidateInputNotNil(Booking.IdCustomer, Booking.CheckInDeadline, Booking.IdRoom)
	if err != nil {
		return err
	}
	err = s.BookingRepo.CreateBooking(Booking)
	if err != nil {
		return err
	}

	return nil
}
