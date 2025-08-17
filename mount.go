package api2

// Mount holds an md.Mount for use in the http server
type Mount struct {
	ID         string                 `json:"id"`
	Path       string                 `json:"path"`
	FileSystem string                 `json:"filesystem,omitempty"`
	Specials   *SpecialPerms          `json:"specials,omitempty"`
	Princs     map[string]*MountPerms `json:"princs,omitempty"`
	FSChain    []*FileSystem          `json:"fsChain,omitempty"`
}

func (mnt *Mount) Createable() error {
	switch {
	case mnt.ID != "":
		return ErrCannotSetID
	case mnt.FSChain != nil:
		return ErrCannotCreateFSChain
	}

	return nil
}

func (mnt *Mount) Updateable() error {
	switch {
	case mnt.ID != "":
		return ErrCannotUpdateID
	case mnt.Path != "":
		return ErrCannotUpdateMountPath
	case mnt.FSChain != nil:
		return ErrCannotUpdateMountFSChain
	}

	return nil
}

// MountPerms holds permissions for a princ with respect to a mount
type MountPerms struct {
	Perms     Perm  `json:"perms,omitempty"`
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
