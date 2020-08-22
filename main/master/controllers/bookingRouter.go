package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/disebud/reservation-hotel/main/master/models"
	"github.com/disebud/reservation-hotel/main/master/usecases"
	"github.com/disebud/reservation-hotel/main/master/utils"
	"github.com/gorilla/mux"
)

type BookingHandler struct {
	BookingUseCase usecases.BookingUseCase
}

func BookingController(r *mux.Router, service usecases.BookingUseCase) {
	BookingHandler := BookingHandler{service}
	r.HandleFunc("/Booking/info", BookingHandler.ListInfoRoom).Methods(http.MethodGet)
	r.HandleFunc("/Booking", BookingHandler.CreateBooking).Methods(http.MethodPost)

}

func (e BookingHandler) ListInfoRoom(w http.ResponseWriter, r *http.Request) {
	currentDate := time.Now()
	Bookings, err := e.BookingUseCase.GetAllBookings()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	response := utils.Response{}
	response.Status = http.StatusOK
	response.Message = "All Data Info Room Booking"
	response.Date = "Today / " + currentDate.Format("02-January-2006")
	response.Result = Bookings
	byteOfBookings, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfBookings))
}

func (e BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var re models.Booking
	err := json.NewDecoder(r.Body).Decode(&re)
	if err != nil {
		log.Println(err)
	}

	err = e.BookingUseCase.CreateBooking(re)
	if err != nil {
		log.Println(err)
	}
	log.Println("Insert successful")
	w.Write([]byte("Insert successful"))

}
