package api2

// AuthInput is for authenticating with an efmrl
type AuthInput struct {
	Value string `json:"value"`
}

// NoSessionCookie is given when there is no session cookie, or when the session
// cookie is corrupt
type NoSessionCookie struct {
	Message string `json:"message"`
}

// SessionNotFound is given when a session ID is found but the session is not
type SessionNotFound struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

// EmptySessionRequest is sent when requesting a new empty session cookie. The
// CookieOK field must be true, or no cookie will be created.
type EmptySessionRequest struct {
	CookieOK bool `json:"cookie_ok"`
}
