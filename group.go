package api2

type PostGroupReq struct {
	Name string `json:"name"`
}

type GetGroupsRes struct {
	Groups []*Group `json:"groups"`
}

type PostUserAction int

const (
	PostUserAdd PostUserAction = iota
	PostUserDrop
)

type PostUserToGroupReq struct {
	UserID string         `json:"userID"`
	Action PostUserAction `json:"action"`
}

type GetGroupUsersRes struct {
	Users []*User `json:"users"`
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
	Perms Perm `json:"perms"`

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

	// GroupFlagsClear indicates that the set of flags given should be patched
	// to be the new value for the group's flags. This is necessary to
	// differentiate whether the flags aren't being modified or if they are
	// being set to all-zero. If present, it must be the only flags set.
	GroupFlagsClear

	// GroupFlagsBeyond is the first undefined flag
	GroupFlagsBeyond

	// GroupFlagsDefined is the set of valid flags
	GroupFlagsDefined = GroupFlagsBeyond - 1
)

// GroupIDs is a list of group IDs
type GroupIDs []string
