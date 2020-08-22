package utils

// Response Struct
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Date    string      `json:"date"`
	Result  interface{} `json:"result"`
}

func GenerateResponse(status int, message, date string, result interface{}) Response {
	var responseInfo Response
	responseInfo.Status = status
	responseInfo.Message = message
	responseInfo.Date = date
	responseInfo.Result = result
	return responseInfo
}
