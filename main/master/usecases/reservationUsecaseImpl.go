package usecases

import (
	"github.com/disebud/reservation-hotel/main/master/models"
	"github.com/disebud/reservation-hotel/main/master/repositories"
)

type ReservationUsecaseImpl struct {
	ReservationRepo repositories.ReservationRepository
}

func InitReservationUsecase(ReservationRepo repositories.ReservationRepository) ReservationUseCase {
	return &ReservationUsecaseImpl{ReservationRepo}
}

func (s ReservationUsecaseImpl) GetAllReservationsPagination(orderBy, sort, page, limit string) ([]*models.Room, error) {
	Reservation, err := s.ReservationRepo.GetAllReservationPagination(orderBy, sort, page, limit)
	if err != nil {
		return nil, err
	}

	return Reservation, nil
}

func (s ReservationUsecaseImpl) GetAllReservations() ([]*models.Room, error) {
	Reservation, err := s.ReservationRepo.GetAllReservation()
	if err != nil {
		return nil, err
	}

	return Reservation, nil
}

func (s ReservationUsecaseImpl) GetReservationByIdRoom(IdRoom string) (*models.Room, error) {
	Reservation, err := s.ReservationRepo.GetReservationByIdRoom(IdRoom)
	if err != nil {
		return nil, err
	}
	return Reservation, nil
}
func (s ReservationUsecaseImpl) GetReservationByStatus(StatusKey string) ([]*models.Room, error) {
	Reservation, err := s.ReservationRepo.GetReservationByStatus(StatusKey)
	if err != nil {
		return nil, err
	}
	return Reservation, nil
}

func (s ReservationUsecaseImpl) DeleteReservationByIdRoom(IdRoom string) (*models.Room, error) {
	Reservation, err := s.ReservationRepo.DeleteReservationByIdRoom(IdRoom)
	if err != nil {
		return nil, err
	}
	return Reservation, nil
}

func (s ReservationUsecaseImpl) CreateReservation(Reservation models.Room) error {
	// err := utils.ValidateInputLenCharacter(3, 6, Reservation.IdRoom)
	// if err != nil {
	// 	return err
	// }
	// err1 := utils.ValidateInputNotNil(Reservation.NameRoom, Reservation.Location, Reservation.Price)
	// if err1 != nil {
	// 	return err
	// }
	err := s.ReservationRepo.CreateReservation(Reservation)
	if err != nil {
		return err
	}

	return nil
}

func (s ReservationUsecaseImpl) UpdateInfoReservationRoom(IdReservation string, Reservation models.Room) error {
	err := s.ReservationRepo.UpdateInfoReservationRoom(IdReservation, Reservation)
	if err != nil {
		return err
	}
	return nil
}

// func (s ReservationUsecaseImpl) GetCountReservationType(StatusKey string) ([]*models.ReservationType, error) {
// 	Reservation, err := s.ReservationRepo.GetCountReservationType(StatusKey)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return Reservation, nil
// }
