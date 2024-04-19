package api2

type PostGroupReq struct {
	Name string `json:"name"`
}

type GetGroupsRes struct {
	Groups []*Group `json:"groups"`
}

// Group defines a group within an efmrl
type Group struct {
	ID      string `json:"id"`
	Created string `json:"created"`
	Name    string `json:"name"`

	// Managers is the ID of a group whose users can manage membership in this
	// group
	Managers string `json:"managers,omitempty"`

	// Perms is a set of permissions that are added to members of this group
	Perms Perm `json:"perms,omitempty"`

	// Flags is a set of group-only booleans; see Group* constants below
	Flags uint64 `json:"flags,omitempty"`
}

const (
	// GroupVisible indicates that this group will be returned when retrieving
	// all groups, even if the user would not have permission otherwise.
	GroupVisible = (1 << iota)

	// GroupUsersVisible indicates that members of this group will be able to
	// see all users within the group. Not all user metadata will be visible,
	// e.g. not the creds.
	GroupUsersVisible

	// GroupJoinable indicates that any user may join this group without being
	// permitted otherwise
	GroupJoinable
)

// GroupIDs is a list of group IDs
type GroupIDs []string
