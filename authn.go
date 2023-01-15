package api2

// SessionReq is sent when requesting a new empty session cookie. The
// CookieOK field must be true, or no cookie will be created.
type SessionReq struct {
	CookieOK bool `json:"cookie_ok"`
}

// SessionRes is returned with the request
type SessionRes struct {
	Message string `json:"message,omitempty"`
	Created string `json:"created,omitempty"`
}
