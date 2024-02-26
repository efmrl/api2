package api2

type WANFinishRegisterRes struct {
	// Message can hold success or error messages
	Message string `json:"message"`
}

type WANFinishLoginRes struct {
	// Message can hold success or error messages
	Message string `json:"message"`
}

// NewSessionReq is sent to the REST API '/newsession' endpoint
type NewSessionReq struct {
	// CookieOK indicates that the user is okay with cookies
	CookieOK bool `json:"cookie_ok"`

	// UserKey specifies a user that we would like to authenticate against
	UserKey string `json:"user_key,omitempty"`

	// UserSecret is a secret given to the user, to finish a login. Usually it's
	// a random string sent via email.
	UserSecret string `json:"user_secret,omitempty"`
}

// NewSessionRes is returned from session requests
type NewSessionRes struct {
	// ID is the ID of the session
	ID string `json:"id,omitempty"`

	// Created is when the session was created, in RFC 3339 format.
	Created string `json:"created,omitempty"`

	// UserKey reflects the UserKey that was set in the request, if it was
	// present.
	UserKey string `json:"user_key,omitempty"`

	// UserID returns a user ID for the user that was specified in the
	// KserKey field of NewSessionReq.
	UserID string `json:"user,omitempty"`

	// Declared is when a user was declared, in RFC 3339 format. Will be empty
	// if no user has been declared.
	Declared string `json:"declared,omitempty"`

	// Confirmed is when a user was confirmed, in RFC 3339 format. Will be empty
	// if no user has been confirmed.
	Confirmed string `json:"confirmed,omitempty"`
}

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
	Creds map[string]*SessionLogin `json:"creds,omitempty"`
	// Users is a list of the IDs of all users that this session has access to.
	Users map[string]*UserLogin `json:"users,omitempty"`
}

type SessionDeleteRes struct {
	Message string `json:"message,omitempty"`
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

// SessionLogin describes a successful login of a cred
type SessionLogin struct {
	When     string `json:"when"`
	Friendly string `json:"friendly"`
}

// UserLogin describes a successful login of a user
type UserLogin struct {
	When     string `json:"when"`
	Friendly string `json:"friendly"`
}
