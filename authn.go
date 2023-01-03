package api2

// AuthReq is for authenticating with an efmrl
type AuthReq struct {
	Value string `json:"value"`
}

// AuthRes is returned after an authentication attempt
type AuthRes struct {
	Message string `json:"message"`
}

// EmptySessionReq is sent when requesting a new empty session cookie. The
// CookieOK field must be true, or no cookie will be created.
type EmptySessionReq struct {
	CookieOK bool `json:"cookie_ok"`
}

// EmptySessionErr is returned with the request
type EmptySessionErr struct {
	Message string `json:"message"`
}
