package api2

// Group defines a group within an efmrl
type Group struct {
	ID      string `json:"id"`
	Created string `json:"created"`
	Name    string `json:"name"`
	Perms   Perm   `json:"perms,omitempty"`
}

// GroupIDs is a list of group IDs
type GroupIDs []string
