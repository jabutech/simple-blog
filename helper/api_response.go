package helper

type ResponseFailed struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Fields  interface{} `json:"fields"`
}

func ApiResponseFailed(code int, status string, message string, fields interface{}) ResponseFailed {
	jsonResponse := ResponseFailed{
		Code:    code,
		Status:  status,
		Message: message,
		Fields:  fields,
	}

	return jsonResponse
}

type ResponseSuccess struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func ApiResponseSuccess(code int, status string, message string) ResponseSuccess {
	jsonResponse := ResponseSuccess{
		Code:    code,
		Status:  status,
		Message: message,
	}

	return jsonResponse
}
