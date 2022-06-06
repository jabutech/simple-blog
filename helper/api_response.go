package helper

type ResponseWithData struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ApiResponseWithData(code int, status string, message string, data interface{}) ResponseWithData {
	jsonResponse := ResponseWithData{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}

	return jsonResponse
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ApiResponseWithoutData(code int, status string, message string) ResponseWithoutData {
	jsonResponse := ResponseWithoutData{
		Code:    code,
		Status:  status,
		Message: message,
	}

	return jsonResponse
}
