package repositories

import "github.com/disebud/reservation-hotel/main/master/models"

type BookingRepository interface {
	CreateBooking(Booking models.Booking) error
	GetAllBooking() ([]*models.Booking, error)
}
