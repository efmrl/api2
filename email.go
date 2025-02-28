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
	// Validated is true if the email has been validated by its user
	Validated string `json:"validated,omitempty"`
}

// CredIDs is a bunch of cred IDs
type CredIDs []string
