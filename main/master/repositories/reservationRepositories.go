package repositories

import (
	"database/sql"

	"github.com/disebud/reservation-hotel/main/master/models"
)

type ReservationRepoImpl struct {
	db *sql.DB
}

func InitReservationRepoImpl(db *sql.DB) ReservationRepository {
	return &ReservationRepoImpl{db}
}

func (s ReservationRepoImpl) GetAllReservation() ([]*models.Room, error) {
	dataReservationRoom := []*models.Room{}
	query := `SELECT mr.id_room , mr.name_room,mr.floor_location,mr.price,mrs.name_status_room
	FROM
		m_room mr
			JOIN
		m_status_room mrs ON mr.status_room = mrs.id_status_room  ORDER BY mr.id_room ASC;`
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		Reservation := models.Room{}
		var err = data.Scan(&Reservation.IdRoom, &Reservation.NameRoom, &Reservation.Location, &Reservation.Price, &Reservation.Status)
		if err != nil {
			return nil, err
		}
		dataReservationRoom = append(dataReservationRoom, &Reservation)
	}

	return dataReservationRoom, nil
}

func (s *ReservationRepoImpl) CreateReservation(Reservation models.Room) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO m_room VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(Reservation.IdRoom, Reservation.NameRoom, Reservation.Location, Reservation.Price, Reservation.Status)
	if err != nil {
		return tx.Rollback()
	}
	stmt.Close()
	return tx.Commit()
}

func (s ReservationRepoImpl) GetReservationByIdRoom(ReservationIdRoom string) (*models.Room, error) {
	Reservation := new(models.Room)
	query := `SELECT mr.id_room , mr.name_room,mr.floor_location,mr.price,mrs.name_status_room
	FROM
		m_room mr
			JOIN
		m_status_room mrs ON mr.status_room = mrs.id_status_room  WHERE mr.id_room = ?;`
	if err := s.db.QueryRow(query, ReservationIdRoom).Scan(&Reservation.IdRoom, &Reservation.NameRoom, &Reservation.Location, &Reservation.Price, &Reservation.Status); err != nil {
		return nil, err
	}
	return Reservation, nil
}

func (s ReservationRepoImpl) GetReservationByStatus(StatusKey string) ([]*models.Room, error) {
	// Reservation := new(models.Reservation)
	dataReservation := []*models.Room{}
	query := `SELECT mr.id_room , mr.name_room,mr.floor_location,mr.price,mrs.id_status_room,mrs.name_status_room
	FROM
		m_room mr
			JOIN
		m_status_room mrs ON mr.status_room = mrs.id_status_room WHERE mrs.id_status_room = ?`
	data, err := s.db.Query(query, StatusKey)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		Reservation := models.Room{}
		var err = data.Scan(&Reservation.IdRoom, &Reservation.NameRoom, &Reservation.Location, &Reservation.Price, &Reservation.IdStatus, &Reservation.Status)
		if err != nil {
			return nil, err
		}
		dataReservation = append(dataReservation, &Reservation)
	}

	return dataReservation, nil
}

func (s ReservationRepoImpl) DeleteReservationByIdRoom(ReservationIdRoom string) (*models.Room, error) {
	Reservation := new(models.Room)
	query := "DELETE FROM m_room WHERE id_room = ?"
	_, err := s.db.Query(query, ReservationIdRoom)
	if err != nil {
		return nil, err
	}
	return Reservation, nil

}

func (s *ReservationRepoImpl) UpdateInfoReservationRoom(id string, Reservation models.Room) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`UPDATE m_room SET name_room = ?,floor_locatiion=?, price = ?, status_room = ? WHERE id_room = ?`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(Reservation.NameRoom, Reservation.Location, Reservation.Price, Reservation.Status, Reservation.IdRoom)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (s *ReservationRepoImpl) UpdateReservationStatusRoom(id string, Reservation models.Room) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`UPDATE m_room SET status_room = ? WHERE id_room = ?`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(Reservation.NameRoom, Reservation.Location, Reservation.Price, Reservation.Status, Reservation.IdRoom)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

// func (s ReservationRepoImpl) GetCountReservationType(StatusKey string) ([]*models.ReservationType, error) {
// 	// Reservation := new(models.Reservation)
// 	dataReservation := []*models.ReservationType{}
// 	query := `SELECT type_Reservation,count(type_Reservation) as Total_Status
// 	FROM
// 		tb_Reservation WHERE type_Reservation= ? AND status_Reservation="ACTIVE";`
// 	data, err := s.db.Query(query, StatusKey)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for data.Next() {
// 		Reservation := models.ReservationType{}
// 		var err = data.Scan(&Reservation.StatusReservation, &Reservation.Total)
// 		if err != nil {
// 			return nil, err
// 		}
// 		dataReservation = append(dataReservation, &Reservation)
// 	}

// 	return dataReservation, nil
// }
