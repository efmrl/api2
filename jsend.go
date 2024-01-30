package api2

// api2 wants to follow JSend: https://github.com/omniti-labs/jsend

const (
	StatusSuccess = "success"
	StatusFail    = "fail"
	StatusError   = "error"
)

type Response struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewSuccess(payload any) *Response {
	return &Response{
		Status: StatusSuccess,
		Data:   payload,
	}
}

func NewFailure(message string) *Response {
	return &Response{
		Status:  StatusFail,
		Message: message,
	}
}

func NewError(message string) *Response {
	return &Response{
		Status:  StatusError,
		Message: message,
	}
}
