package api2

// api2 wants to follow JSend: https://github.com/omniti-labs/jsend
// found from https://stackoverflow.com/questions/12806386/is-there-any-standard-for-json-api-response-format

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

type APIer interface {
	API() any
}

func NewSuccess(payload APIer) *Response {
	return &Response{
		Status: StatusSuccess,
		Data:   payload.API(),
	}
}

func NewSuccessAny(payload any) *Response {
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
