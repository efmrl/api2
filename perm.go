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

	// PermReadMounts gives permission to read the mount table
	PermReadMounts
	// PermUpdateMounts gives permission to update existing mounts
	PermUpdateMounts

	// PermUndefined means "undefined", and it's used in looping. It MUST be
	// the last of the (1 << iota) assignments.
	PermUndefined

	// PermAllDefined is dumb
	PermAllDefined = PermUndefined - 1

	// PermAll is all perms, including currently invalid ones. This is used for
	// "future-proofing" for efmrl owners, as more permission bits are defined.
	PermAll = Perm(0xffffffffffffffff)

	// PermFiles comprises the permission flags that relate to the file system
	PermFiles = PermRead | PermOverwrite | PermCreate | PermDelete

	// PermReadOnly comprises the permission flags that are "read-only"
	// operations: operations that don't affect any data or metadata, etc.
	PermReadOnly = PermRead | PermReadMounts
)

// SpecialPerms holds perms for special princs
type SpecialPerms struct {
	Everyone      Perm `json:"everyone"`
	Sessioned     Perm `json:"sessioned"`
	Authenticated Perm `json:"authenticated"`
}
