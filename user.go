package api2

// User defines a user within an efmrl
type User struct {
	ID      string   `json:"id"`
	Created string   `json:"created"`
	Name    string   `json:"name"`
	Creds   CredIDs  `json:"creds,omitempty"`
	Perms   int      `json:"perms,omitempty"`
	Groups  GroupIDs `json:"groups,omitempty"`
}

// UserIDs is a bunch of user IDs
type UserIDs []string
