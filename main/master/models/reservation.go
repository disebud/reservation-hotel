package models

type Room struct {
	IdRoom   string `json:"id_room"`
	NameRoom string `json:"name_room"`
	Location string `json:"location"`
	Price    int    `json:"price"`
	IdStatus string `json:"id_status"`
	Status   string `json:"status"`
}

type Booking struct {
	CodeBooking     string `json:"code_booking"`
	IdCustomer      string `json:"id_customer"`
	DateBooking     string `json:"date_booking"`
	CheckInDeadline string `json:"check_in_deadline"`
	IdRoom          string `json:"id_room"`
	CodePayment     string `json:"code_payment"`
}

type Customer struct {
	IdCustomer   string `json:"id_customer"`
	NameCustomer string `json:"name_customer"`
	Address      string `json:"address"`
	Telephone    int    `json:"telephone"`
	Email        string `json:"email"`
}

type Rooms []Room
type TotalRoom struct {
	TotalItem int   `json:"totalitem"`
	Room      Rooms `json:"room"`
}
