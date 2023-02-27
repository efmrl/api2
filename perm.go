package api2

// Perm represents permissions
type Perm uint64

const (
	// PermNone is no perms
	PermNone = Perm(0)

	// PermFirst is used in looping
	PermFirst Perm = Perm(1 << iota)

	// PermRead gives read permission
	PermRead
	// PermOverwrite gives permission to write over an existing file
	PermOverwrite
	// PermCreate gives permission to create a new file
	PermCreate
	// PermDelete gives permission to delete a file
	PermDelete

	// PermUndefined means "undefined", and it's used in looping
	PermUndefined

	// PermAll is all valid perms
	PermAll = PermUndefined - 1

	// PermFiles comprises the permission flags that relate to the file system
	PermFiles = PermRead | PermOverwrite | PermCreate | PermDelete

	// PermReadOnly comprises the permission flags that are "read-only"
	// operations: operations that don't affect any data or metadata, etc.
	PermReadOnly = PermRead
)

// SpecialPerms holds perms for special princs
type SpecialPerms struct {
	Everyone      Perm `json:"everyone"`
	Sessioned     Perm `json:"sessioned"`
	Authenticated Perm `json:"authenticated"`
}
