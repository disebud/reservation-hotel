package usecases

import "github.com/disebud/reservation-hotel/main/master/models"

type ReservationUseCase interface {
	GetAllReservations() ([]*models.Room, error)
	CreateReservation(Reservation models.Room) error
	GetReservationByIdRoom(ReservationIdRoom string) (*models.Room, error)
	GetReservationByStatus(StatusKey string) ([]*models.Room, error)
	DeleteReservationByIdRoom(ReservationIdRoom string) (*models.Room, error)
	UpdateInfoReservationRoom(id string, Reservation models.Room) error
}
