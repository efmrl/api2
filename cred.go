package api2

const (
	// CredTypeMulti designates a cred as a Multi (see below)
	CredTypeMulti = "multi"
	// CredTypeEmail designates a cred as an Email (see below)
	CredTypeEmail = "email"
)

// Multi is a type of cred that consists of multiple sub-creds.
type Multi struct {
	ID string `json:"id"`
	// Type will always be CredTypeMulti, i.e. "multi"
	Type string `json:"type"`
	// Children is all of the sub-creds for this cred
	Children CredIDs `json:"children"`
	// Needed is the number of child creds which must be in your session to
	// be considered to have this cred. A value of 0 means that all child creds
	// must be in your session to qualify.
	Needed int `json:"needed"`
}

// Email is a type of cred that represents an email address.
type Email struct {
	ID string `json:"id"`
	// Type will always be CredTypeEmail, i.e. "email"
	Type string `json:"type"`
	// Address is the email address
	Address string `json:"address"`
	// Name is the name of this cred, e.g. "work" for a work email address
	Name string `json:"name,omitempty"`
}

// CredIDs is a bunch of cred IDs
type CredIDs []string
