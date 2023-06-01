package api2

// SessionReq is sent to the REST API '/session' endpoint.
type SessionReq struct {
	// CookieOK must be true for a new session cookie to be issued.
	CookieOK bool `json:"cookie_ok"`
	// Input, if given, is the input for an authentication attempt.
	Input string `json:"input,omitempty"`
}

// SessionRes is returned with the request
type SessionRes struct {
	// Message can hold success or error messages
	Message string `json:"message,omitempty"`

	// ID is the session ID
	ID string `json:"id,omitempty"`
	// Created is when the session was created, in RFC 3339 format.
	Created string `json:"created,omitempty"`
	// Creds is a list of the IDs of all credentials that this session holds.
	Creds CredIDs `json:"creds,omitempty"`
	// Users is a list of the IDs of all users that this session has access to.
	Users UserIDs `json:"users,omitempty"`
}

// SessionEfmrlsReq struct {
type SessionEfmrlsReq struct {
	StartAfter string `json:"start_after,omitempty"`

	MaxCount int32 `json:"max_count,omitempty"`
}

// SessionEfmrlsRes struct {
type SessionEfmrlsRes struct {
	Message string `json:"message,omitempty"`

	IDs []string `json:"ids"`

	ExclNext string `json:"excl_next,omitempty"`

	Count int32 `json:"count"`
}
