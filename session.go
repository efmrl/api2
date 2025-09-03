package api2

// SessionReq is sent to the REST API '/session' endpoint
type SessionReq struct {
	// CookieOK indicates that the user is okay with cookies
	CookieOK bool `json:"cookie_ok"`

	// UserKey specifies a user that we would like to authenticate against
	UserKey string `json:"user_key,omitempty"`

	// UserSecret is a secret given to the user, to finish a login. Usually it's
	// a random string sent via email.
	UserSecret string `json:"user_secret,omitempty"`
}

// SessionRes is returned from session requests
type SessionRes struct {
	// ID is the ID of the session
	ID string `json:"id,omitempty"`

	// Created is when the session was created, in RFC 3339 format.
	Created string `json:"created,omitempty"`

	// UserKey reflects the UserKey that was set in the request, if it was
	// present.
	UserKey string `json:"user_key,omitempty"`

	// UserID returns a user ID for the user that was specified in the
	// UserKey field of SessionReq.
	UserID string `json:"user,omitempty"`

	// Declared is when a user was declared, in RFC 3339 format. Will be empty
	// if no user has been declared.
	Declared string `json:"declared,omitempty"`

	// Confirmed is when a user was confirmed, in RFC 3339 format. Will be empty
	// if no user has been confirmed.
	Confirmed string `json:"confirmed,omitempty"`

	// Active is when the session has last been used. It is imperfect, but
	// it roughly corresponds to when the session came into play when testing
	// for permission or identity.
	Active string `json:"active,omitempty"`

	// UserName is the user.name field
	UserName string `json:"userName,omitempty"`

	// Perms is the current set of permissions in this session
	Perms Perm `json:"perms,omitempty"`
}

type SessionsRes struct {
	// Sessions is the list of active sessions.
	Sessions []*SessionRes `json:"sessions"`
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

// GetTokenRes returns a login token
type GetTokenRes struct {
	// Token is the login token
	Token string `json:"token"`
	// Expires is when the token expires, in RFC 3339 format.
	Expires string `json:"expires,omitempty"`
}
