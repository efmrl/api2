package api2

// PostUserReq is used to request a user to be created
type PostUserReq struct {
	User  *User  `json:"user"`
	Email *Email `json:"email,omitempty"`
}

// PostUserRes is returned after a user is posted
type PostUserRes struct {
	Message string `json:"message,omitempty"`
	User    *User  `json:"user,omitempty"`
}

// PatchUserReq is used to patch a user's metadata
type PatchUserReq struct {
	User *User `json:"user"`
}

// PatchUserRes is returned after a user is patched
type PatchUserRes struct {
	Message string `json:"message,omitempty"`
	User    *User  `json:"user"`
}

// GetUsersRes is returned for a GET on users
type GetUsersRes struct {
	Message string  `json:"message,omitempty"`
	Users   []*User `json:"users"`
}

// User defines a user within an efmrl
type User struct {
	ID      string   `json:"id"`
	Created string   `json:"created"`
	Name    string   `json:"name"`
	Creds   CredIDs  `json:"creds,omitempty"`
	Perms   Perm     `json:"perms,omitempty"`
	Groups  GroupIDs `json:"groups,omitempty"`
}

// UserIDs is a bunch of user IDs
type UserIDs []string
