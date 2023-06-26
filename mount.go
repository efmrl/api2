package api2

// Mount holds an md.Mount for use in the http server
type Mount struct {
	FileSystem string                 `json:"filesystem,omitempty"`
	Specials   SpecialPerms           `json:"specials,omitempty"`
	Princs     map[string]*MountPerms `json:"princs,omitempty"`
}

// MountPerms holds permissions for a princ with respect to a mount
type MountPerms struct {
	Perms     Perm  `json:"perms"`
	FileQuota int32 `json:"file_quota"`
	ByteQuota int32 `json:"byte_quota"`
}

// GetMountsRes is the result of getting an efmrl's mounts
type GetMountsRes struct {
	// Message can hold success or error messages
	Message string `json:"message"`

	// Mounts is the list of all mounts on the efmrl
	Mounts map[string]*Mount `json:"mounts"`
}

// PostMountReq is a request to create a mount
type PostMountReq struct {
	Path  string `json:"path"`
	Mount Mount  `json:"mount"`
}

// PostMountRes is the result of posting a mount
type PostMountRes struct {
	Message string `json:"message"`
}

// DeleteMountRes is the result of deleting a mount
type DeleteMountRes struct {
	Message string `json:"message"`
}
