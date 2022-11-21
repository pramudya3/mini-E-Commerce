package helpers

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseSuccessWithoutData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseFailed struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func APIResponseSuccess(code int, message string, data interface{}) ResponseSuccess {
	jsonResponse := ResponseSuccess{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return jsonResponse
}

func APIResponseSuccessWithoutData(code int, message string) ResponseSuccessWithoutData {
	jsonResponse := ResponseSuccessWithoutData{
		Code:    code,
		Message: message,
	}
	return jsonResponse
}

func APIResponseFailed(code int, message string) ResponseFailed {
	jsonResponse := ResponseFailed{
		Code:    code,
		Message: message,
	}
	return jsonResponse
}
