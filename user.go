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

// PutUserReq is used to put a user's metadata
type PutUserReq struct {
	User *User `json:"user"`
}

// PutUserRes is returned after a user is put
type PutUserRes struct {
	Message string `json:"message,omitempty"`
	User    *User  `json:"user"`
}

// GetUserRes is returned for a GET on a particular user
type GetUserRes struct {
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
	Perms   Perm     `json:"perms,omitempty"`
	Groups  GroupIDs `json:"groups,omitempty"`
}
