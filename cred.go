package api2

const (
	// CredTypeMulti designates a cred as a Multi (see below)
	CredTypeMulti = "multi"
	// CredTypeEmail designates a cred as an Email (see below)
	CredTypeEmail = "email"
	// CredTypeTOTP designates a cred as a TOTP (see below)
	CredTypeTOTP = "totp"
	// CredTypeGithub designates a cred as a login-via-github (see below)
	CredTypeGithub = "github"
)

type GetCredsRes struct {
	Message string    `json:"message,omitempty"`
	Multis  []*Multi  `json:"multis,omitempty"`
	Emails  []*Email  `json:"emails,omitempty"`
	TOTPs   []*TOTP   `json:"totps,omitempty"`
	Githubs []*Github `json:"githubs,omitempty"`
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
	TOTP *TOTP `json:"totp"`
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
}

// Email is a type of cred that represents an email address.
type Email struct {
	ID string `json:"id"`
	// Address is the email address
	Address string `json:"address"`
	// Desc is the description of this cred, e.g. "work" for a work email address
	Desc string `json:"name,omitempty"`
}

// TOTP is a Timed One-Time-Password cred.
type TOTP struct {
	ID string `json:"id"`
	// Desc holds a user-supplied description for this TOTP
	Desc string `json:"desc,omitempty"`
	// number of digits the TOTP will have. Must be 6, 8, or 0 for the default
	// value (6).
	Digits int `json:"digits,omitempty"`
}

// Github is a login-via-github cred.
type Github struct {
	ID          string `json:"id"`
	GithubID    int64  `json:"githubID"`
	GithubLogin string `json:"githubLogin"`
}

// CredIDs is a bunch of cred IDs
type CredIDs []string
