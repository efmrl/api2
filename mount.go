package api2

import "fmt"

var (
	ErrCannotSetID               = fmt.Errorf("ID cannot be set")
	ErrCannotCreateReadOnlyMount = fmt.Errorf("read only mount cannot be specified")
	ErrCannotCreateFSChain       = fmt.Errorf("FSChain cannot be specified")
	ErrCannotUpdateID            = fmt.Errorf("ID cannot be updated")
	ErrCannotUpdateMountPath     = fmt.Errorf("path cannot be updated")
	ErrCannotUpdateMountFSChain  = fmt.Errorf("FSChain cannot be updated")
)

// Mount holds an md.Mount for use in the http server
type Mount struct {
	ID         string                 `json:"id"`
	Path       string                 `json:"path"`
	FileSystem string                 `json:"filesystem,omitempty"`
	ReadOnly   bool                   `json:"readOnly,omitempty"`
	Specials   *SpecialPerms          `json:"specials,omitempty"`
	Princs     map[string]*MountPerms `json:"princs,omitempty"`
	FSChain    []*FileSystem          `json:"fsChain,omitempty"`
}

func (mnt *Mount) Createable() error {
	switch {
	case mnt.ID != "":
		return ErrCannotSetID
	case mnt.ReadOnly:
		return ErrCannotCreateReadOnlyMount
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
