package response

type Response struct {
	Code       uint64      `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	ErrorField interface{} `json:"errorField,omitempty"`
	StatusCode int         `json:"statusCode"`
}

func OK(i interface{}) *Response {
	return &Response{
		Code:       0,
		Message:    "success",
		Data:       i,
		StatusCode: 200,
	}
}

func Err(code uint64, statusCode int, message string) *Response {
	return &Response{
		Code:       code,
		StatusCode: statusCode,
		Message:    message,
	}
}

func ErrField(code uint64, statusCode int, message string, errorField interface{}) *Response {
	return &Response{
		Code:       code,
		Message:    message,
		ErrorField: errorField,
		StatusCode: statusCode,
	}
}
