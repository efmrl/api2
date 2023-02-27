package api2

// Mount holds an md.Mount for use in the http server
type Mount struct {
	FileSystem string                 `json:"filesystem"`
	Specials   SpecialPerms           `json:"specials"`
	Princs     map[string]*MountPerms `json:"princs"`
}

// MountPerms holds permissions for a princ with respect to a mount
type MountPerms struct {
	Perms     Perm  `json:"perms"`
	FileQuota int32 `json:"file_quota"`
	ByteQuota int32 `json:"byte_quota"`
}
