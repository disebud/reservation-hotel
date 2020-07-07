package repositories

import (
	"database/sql"

	"github.com/disebud/reservation-hotel/main/master/models"
)

type BookingRepoImpl struct {
	db *sql.DB
}

func InitBookingRepoImpl(db *sql.DB) BookingRepository {
	return &BookingRepoImpl{db}
}

func (s *BookingRepoImpl) CreateBooking(Booking models.Booking) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO m_booking VALUES (?, ?, ?, ?, ? , ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(Booking.CodeBooking, Booking.IdCustomer, Booking.DateBooking, Booking.CheckInDeadline, Booking.IdRoom, Booking.CodePayment)
	if err != nil {
		return tx.Rollback()
	}
	stmt.Close()
	return tx.Commit()
}

func (s BookingRepoImpl) GetAllBooking() ([]*models.Booking, error) {
	dataReservationBooking := []*models.Booking{}
	query := `SELECT * FROM m_booking;`
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		Reservation := models.Booking{}
		var err = data.Scan(&Reservation.CodeBooking, &Reservation.IdCustomer, &Reservation.DateBooking, &Reservation.CheckInDeadline, &Reservation.IdRoom, &Reservation.CodePayment)
		if err != nil {
			return nil, err
		}
		dataReservationBooking = append(dataReservationBooking, &Reservation)
	}

	return dataReservationBooking, nil
}
