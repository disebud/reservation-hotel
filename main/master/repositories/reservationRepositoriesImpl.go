package repositories

import "github.com/disebud/reservation-hotel/main/master/models"

type ReservationRepository interface {
	GetAllReservationPagination(orderBy, sort, page, limit string) ([]*models.Room, error)
	GetAllReservation() ([]*models.Room, error)
	// GetAllReservation(page, limit, orderBy, sort string) (*TotalRoom?, error)
	CreateReservation(Reservation models.Room) error
	GetReservationByIdRoom(ReservationIdRoom string) (*models.Room, error)
	GetReservationByStatus(StatusKey string) ([]*models.Room, error)
	DeleteReservationByIdRoom(ReservationIdRoom string) (*models.Room, error)
	UpdateInfoReservationRoom(id string, Reservation models.Room) error
	UpdateReservationStatusRoom(id string, Reservation models.Room) error
}
