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
)
