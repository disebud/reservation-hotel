package repositories

import "github.com/disebud/reservation-hotel/main/master/models"

type ReservationRepository interface {
	GetAllReservation() ([]*models.Room, error)
	CreateReservation(Reservation models.Room) error
	GetReservationByIdRoom(ReservationIdRoom string) (*models.Room, error)
	GetReservationByStatus(StatusKey string) ([]*models.Room, error)
	DeleteReservationByIdRoom(ReservationIdRoom string) (*models.Room, error)
	UpdateInfoReservationRoom(id string, Reservation models.Room) error
	UpdateReservationStatusRoom(id string, Reservation models.Room) error
}
