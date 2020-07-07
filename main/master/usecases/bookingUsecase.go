package usecases

import "github.com/disebud/reservation-hotel/main/master/models"

type BookingUseCase interface {
	GetAllBookings() ([]*models.Booking, error)
	CreateBooking(Reservation models.Booking) error
}
