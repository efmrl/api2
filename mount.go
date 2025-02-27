package api2

// Mount holds an md.Mount for use in the http server
type Mount struct {
	ID         string                 `json:"id"`
	Path       string                 `json:"path"`
	FileSystem string                 `json:"filesystem,omitempty"`
	Specials   *SpecialPerms          `json:"specials,omitempty"`
	Princs     map[string]*MountPerms `json:"princs,omitempty"`
}

// MountPerms holds permissions for a princ with respect to a mount
type MountPerms struct {
	Perms     Perm  `json:"perms"`
	FileQuota int32 `json:"fileQuota"`
	ByteQuota int32 `json:"byteQuota"`
}

// GetMountsRes is the result of getting an efmrl's mounts
type GetMountsRes struct {
	// Mounts is the list of all mounts on the efmrl
	Mounts []*Mount `json:"mounts"`
}

type PostMountRes struct {
	ID string `json:"id"`
}
