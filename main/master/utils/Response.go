package utils

// Response Struct
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Date    string      `json:"date"`
	Data    interface{} `json:"data"`
}
