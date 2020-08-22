package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/disebud/reservation-hotel/main/master/models"
	"github.com/disebud/reservation-hotel/main/master/usecases"
	"github.com/disebud/reservation-hotel/main/master/utils"
	"github.com/gorilla/mux"
)

type ReservationHandler struct {
	ReservationUseCase usecases.ReservationUseCase
}

func ReservationController(r *mux.Router, service usecases.ReservationUseCase) {
	ReservationHandler := ReservationHandler{service}
	r.HandleFunc("/reservation/info", ReservationHandler.ListInfoRoomPagination).Queries("orderBy", "{orderBy}", "sort", "{sort}", "page", "{page}", "limit", "{limit}").Methods(http.MethodGet)
	r.HandleFunc("/reservations", ReservationHandler.ListInfoRoom).Methods(http.MethodGet)
	r.HandleFunc("/reservation/{id}", ReservationHandler.GetReservationIdRoom).Methods(http.MethodGet)
	r.HandleFunc("/reservation/status/{status}", ReservationHandler.GetReservationStatus).Methods(http.MethodGet)
	r.HandleFunc("/reservation", ReservationHandler.CreateReservation).Methods(http.MethodPost)
	r.HandleFunc("/reservation/{id}", ReservationHandler.UpdateReservations).Methods(http.MethodPut)
	r.HandleFunc("/reservation/{id}", ReservationHandler.DeleteReservationID).Methods(http.MethodDelete)
}

func (e ReservationHandler) ListInfoRoomPagination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)

	orderBy, _ := vars["orderBy"]
	sort, _ := vars["sort"]
	page, _ := vars["page"]
	limit, _ := vars["limit"]
	// log.Println(infinfoType)
	currentDate := time.Now()
	Reservations, err := e.ReservationUseCase.GetAllReservationsPagination(orderBy, sort, page, limit)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	response := utils.Response{}
	response.Status = http.StatusOK
	response.Message = "All Data Info Room Reservation"
	response.Date = "Today / " + currentDate.Format("02-January-2006")
	response.Result = Reservations
	byteOfReservations, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfReservations))
}

func (e ReservationHandler) ListInfoRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)

	// log.Println(infinfoType)
	currentDate := time.Now()
	Reservations, err := e.ReservationUseCase.GetAllReservations()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	response := utils.Response{}
	response.Status = http.StatusOK
	response.Message = "All Data Info Room Reservation"
	response.Date = "Today / " + currentDate.Format("02-January-2006")
	response.Result = Reservations
	byteOfReservations, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfReservations))
}

// func (e ReservationHandler) ListReservations(w http.ResponseWriter, r *http.Request) {
// 	currentDate := time.Now()
// 	Reservations, err := e.ReservationUseCase.GetReservations()
// 	if err != nil {
// 		w.Write([]byte("Data Not Found"))
// 	}
// 	response := utils.Response{}
// 	response.Status = http.StatusAccepted
// 	response.Message = "All Data Reservation"
// 	response.Date = "Today / " + currentDate.Format("02-January-2006")
// 	response.Data = Reservations
// 	byteOfReservations, err := json.Marshal(response)
// 	if err != nil {
// 		w.Write([]byte("Oops something when wrong"))
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte(byteOfReservations))
// }

func (e ReservationHandler) GetReservationIdRoom(w http.ResponseWriter, r *http.Request) {
	currentDate := time.Now()
	vars := mux.Vars(r)
	Employees, err := e.ReservationUseCase.GetReservationByIdRoom(vars["id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	response := utils.Response{}
	response.Status = http.StatusAccepted
	response.Message = " Data Reservation id Room = " + vars["id"]
	response.Date = "Today / " + currentDate.Format("02-January-2006")
	response.Result = Employees
	byteOfReservations, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops try again"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfReservations)
}

func (e ReservationHandler) GetReservationStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Reservations, err := e.ReservationUseCase.GetReservationByStatus(vars["status"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	response := utils.Response{}
	response.Status = http.StatusAccepted
	response.Message = "All Data Reservation id Status : " + vars["status"]
	response.Result = Reservations
	byteOfReservations, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops try again"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfReservations)
}

func (e ReservationHandler) DeleteReservationID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	_, err := e.ReservationUseCase.DeleteReservationByIdRoom(param["id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete Data successful , Room ID = " + param["id"]))
	log.Println("Delete Data successful , Room ID = " + param["id"])
}

func (e ReservationHandler) CreateReservation(w http.ResponseWriter, r *http.Request) {
	var re models.Room
	err := json.NewDecoder(r.Body).Decode(&re)
	if err != nil {
		log.Println(err)
	}
	log.Println(re)
	err = e.ReservationUseCase.CreateReservation(re)
	if err != nil {
		log.Println(err)
	}
	log.Println("Insert successful")
	w.Write([]byte("Insert successful"))

}

func (e ReservationHandler) UpdateReservations(w http.ResponseWriter, r *http.Request) {
	currentDate := time.Now()
	date := "Today / " + currentDate.Format("02-January-2006")
	var reservation models.Room
	param := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&reservation)
	err = e.ReservationUseCase.UpdateInfoReservationRoom(param["id"], reservation)
	if err != nil {
		w.Write([]byte("Data not found !!"))
		log.Println(err)
		return
	}
	byteOfInfo, _ := json.Marshal(utils.GenerateResponse(http.StatusOK, "Success Update", date, reservation))
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfInfo)
	w.Write([]byte("\n" + "Update Data successful , NIK = " + param["id"]))
	log.Println("Update Data successful , NIK = " + param["id"])
}
