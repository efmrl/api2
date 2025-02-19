package api2

// PostUserReq is used to request a user to be created
type PostUserReq struct {
	User  *User  `json:"user"`
	Email *Email `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// ListUsersRes returns all users which can be viewed by the requester
type ListUsersRes struct {
	Users []*User `json:"users"`
}

// User defines a user within an efmrl
type User struct {
	ID      string   `json:"id"`
	Created string   `json:"created"`
	Name    string   `json:"name"`
	Perms   Perm     `json:"perms"`
	Emails  []*Email `json:"emails,omitempty"`
}
