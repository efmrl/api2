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

// NewResult gives you a target for JSON decoding. Pass in the structure to be
// filled in on success.
func NewResult(success any) *Response {
	return &Response{
		Data: success,
	}
}

func NewSuccess(payload APIer) *Response {
	return &Response{
		Status: StatusSuccess,
		Data:   payload.API(),
	}
}

func NewSuccessMessage(payload APIer, message string) *Response {
	res := NewSuccess(payload)
	res.Message = message
	return res
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

func NewError(err error) *Response {
	return &Response{
		Status:  StatusError,
		Message: err.Error(),
	}
}
