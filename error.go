package api2

import "fmt"

var (
	ErrCannotSetID               = fmt.Errorf("ID cannot be set")
	ErrCannotSetCreate           = fmt.Errorf("create time cannot be set")
	ErrCannotCreateReadOnlyMount = fmt.Errorf("read only mount cannot be specified")
	ErrCannotCreateFSChain       = fmt.Errorf("FSChain cannot be specified")
	ErrCannotUpdateID            = fmt.Errorf("ID cannot be updated")
	ErrCannotUpdateMountPath     = fmt.Errorf("path cannot be updated")
	ErrCannotUpdateMountFSChain  = fmt.Errorf("FSChain cannot be updated")
)
