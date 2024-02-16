package api2

type PostEmailReq struct {
	Email *Email `json:"email"`
}

type PostEmailRes struct {
	Message string `json:"message,omitempty"`
	Email   *Email `json:"email,omitempty"`
}

type PutEmailReq struct {
	Email *Email `json:"email"`
}

type PutEmailRes struct {
	Message string `json:"message,omitempty"`
	Email   *Email `json:"email,omitempty"`
}

type DeleteCredRes struct {
	Message string `json:"message,omitempty"`
}

// Email is a type of cred that represents an email address.
type Email struct {
	ID string `json:"id"`
	// Address is the email address
	Address string `json:"address"`
	// Desc is the description of this cred, e.g. "work" for a work email address
	Desc string `json:"name,omitempty"`
	// UserID is the id of the user that owns this cred
	UserID string `json:"user_id"`
	// InSession is true if this cred is in the current session
	InSession bool `json:"in_session"`
}

// CredIDs is a bunch of cred IDs
type CredIDs []string
