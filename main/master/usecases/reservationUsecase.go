package usecases

import "github.com/disebud/reservation-hotel/main/master/models"

type ReservationUseCase interface {
	GetAllReservationsPagination(orderBy, sort, page, limit string) ([]*models.Room, error)
	GetAllReservations() ([]*models.Room, error)
	CreateReservation(Reservation models.Room) error
	GetReservationByIdRoom(ReservationIdRoom string) (*models.Room, error)
	GetReservationByStatus(StatusKey string) ([]*models.Room, error)
	DeleteReservationByIdRoom(ReservationIdRoom string) (*models.Room, error)
	UpdateInfoReservationRoom(id string, Reservation models.Room) error
}
