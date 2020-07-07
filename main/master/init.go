package master

import (
	"database/sql"

	"github.com/disebud/reservation-hotel/main/master/controllers"
	"github.com/disebud/reservation-hotel/main/master/repositories"
	"github.com/disebud/reservation-hotel/main/master/usecases"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	reservationRepo := repositories.InitReservationRepoImpl(db)
	reservationUseCase := usecases.InitReservationUsecase(reservationRepo)
	controllers.ReservationController(r, reservationUseCase)
	///////////////////////////////////////////////////////
	bookingRepo := repositories.InitBookingRepoImpl(db)
	bookingUseCase := usecases.InitBookingUsecase(bookingRepo)
	controllers.BookingController(r, bookingUseCase)
}
