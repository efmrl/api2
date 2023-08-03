package api2

const (
	// CredTypeMulti designates a cred as a Multi (see below)
	CredTypeMulti = "multi"
	// CredTypeEmail designates a cred as an Email (see below)
	CredTypeEmail = "email"
)

type GetCredsRes struct {
	Message string   `json:"message,omitempty"`
	Multis  []*Multi `json:"multis,omitempty"`
	Emails  []*Email `json:"emails,omitempty"`
}

type PostMultiReq struct {
	Multi *Multi `json:"multi"`
}

type PostMultiRes struct {
	Message string `json:"message,omitempty"`
	Multi   *Multi `json:"multi,omitempty"`
}

type PostEmailReq struct {
	Email *Email `json:"email"`
}

type PostEmailRes struct {
	Message string `json:"message,omitempty"`
	Email   *Email `json:"email,omitempty"`
}

type PostTOTPReq struct {
	TOTP TOTP `json:"totp"`
}

type PostTOTPRes struct {
	Message string `json:"message,omitempty"`
	TOTP    *TOTP  `json:"totp,omitempty"`
}

// Multi is a type of cred that consists of multiple sub-creds.
type Multi struct {
	ID string `json:"id"`
	// Needed is the number of child creds which must be in your session to
	// be considered to have this cred. A value of 0 means that all child creds
	// must be in your session to qualify.
	Needed int `json:"needed"`
	// Parent is the id of the parent of this cred, if any
	Parent string `json:"parent,omitempty"`
}

// Email is a type of cred that represents an email address.
type Email struct {
	ID string `json:"id"`
	// Address is the email address
	Address string `json:"address"`
	// Desc is the description of this cred, e.g. "work" for a work email address
	Desc string `json:"name,omitempty"`
	// Parent is the id of the parent of this cred, if any
	Parent string `json:"parent,omitempty"`
}

// TOTP is a Timed One-Time-Password cred.
type TOTP struct {
	ID string `json:"id"`
	// Desc holds a user-supplied description for this TOTP
	Desc string `json:"desc,omitempty"`
	// number of digits the TOTP will have. Must be 6, 8, or 0 for the default
	// value (6).
	Digits int `json:"digits,omitempty"`
	// Parent is the id of the parent of this cred
	Parent string `json:"parent,omitempty"`
}

// CredIDs is a bunch of cred IDs
type CredIDs []string
