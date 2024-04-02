package api2

import (
	"fmt"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

// Perm represents permissions
type Perm uint64

const (
	// PermRead gives read permission. It MUST be defined first in this block
	// of constants. Not just the (1 << iota) constants, but everything in this
	// block.
	PermRead = Perm(1 << iota)
	// PermOverwrite gives permission to write over an existing file
	PermOverwrite
	// PermCreate gives permission to create a new file
	PermCreate
	// PermDelete gives permission to delete a file
	PermDelete
	// PermListFiles gives permission to get a listing of file information
	PermListFiles

	// PermReadMounts gives permission to read the mount table
	PermReadMounts
	// PermUpdateMounts gives permission to update existing mounts
	PermUpdateMounts

	// PermCreateUser grants the ability to create a new user
	PermCreateUser

	// PermReadUsers grants the ability to read all users on the efmrl. Every
	// user can read their own metadata.
	PermReadUsers

	// PermWriteUsers grants the ability to create, update, or delete any users.
	// Every user can update their own metadata (other than permissions).
	PermWriteUsers

	// PermReadPerms grants the ability to read permissions on the efmrl,
	// mounts, and all princs which are visible.
	PermReadPerms

	// PermWritePerms grants the ability to edit the permissions on users and
	// groups. Without this permission, permissions cannot be edited. Note that
	// the efmrl's owner always has this permission, even if the permission bits
	// don't reflect it.
	PermWritePerms

	// PermReadGroups grants the ability to read all groups
	PermReadGroups

	// PermEditGroups grants the ability to create, edit, or delete groups; and
	// to add or remove users to groups
	PermWriteGroups

	// PermReadNames grants the ability to read the names given to the efmrl.
	PermReadNames

	// PermWriteNames grants the ability to modify the list of names for the
	// efmrl.
	PermWriteNames

	// PermUndefined means "undefined", and it's used in looping. It MUST be
	// the last of the (1 << iota) assignments.
	PermUndefined

	// PermAllDefined is dumb
	PermAllDefined = PermUndefined - 1

	// PermAll is all perms, including currently invalid ones. This is used for
	// "future-proofing" for efmrl owners, as more permission bits are defined.
	PermAll = ^Perm(0)

	// PermNone is the case where no permissions are given
	PermNone = Perm(0)

	// PermFirst is used in looping
	PermFirst = Perm(1)

	// PermFiles comprises the permission flags that relate to the file system
	PermFiles = PermRead | PermOverwrite | PermCreate | PermDelete |
		PermListFiles

	// PermWrite is the ability to write a file, regardless of whether it
	// currently exists.
	PermWrite = PermOverwrite | PermCreate

	// PermReadOnly comprises the permission flags that are "read-only"
	// operations: operations that don't affect any data or metadata, etc.
	PermReadOnly = PermRead | PermReadMounts | PermReadNames
)

type AllPerms struct {
	Efmrl  *SpecialPerms       `json:"efmrl,omitempty"`
	Mounts map[string]*Mount   `json:"mounts,omitempty"`
	Owner  *NewUser            `json:"owner,omitempty"`
	Users  map[string]*NewUser `json:"users,omitempty"`
	Groups map[string]*Group   `json:"groups,omitempty"`
}

// SpecialPerms holds perms for special princs
type SpecialPerms struct {
	Everyone      *Perm `json:"everyone"`
	Sessioned     *Perm `json:"sessioned"`
	Authenticated *Perm `json:"authenticated"`
}

type PermReference struct {
	StringToBits map[string]Perm `json:"stringToBits"`
	BitsToString map[Perm]string `json:"bitsToString"`
}

// API usually converts an internal-only data structure to a type defined in
// this API. In this case, it simply populates a PermReference, so that a
// caller could use api2.NewSuccess(pr).
func (pr *PermReference) API() any {
	pr.StringToBits = make(map[string]Perm, len(PermNameValue))
	pr.BitsToString = make(map[Perm]string, len(PermNameValue))

	for k, v := range PermNameValue {
		pr.StringToBits[k] = v
		pr.BitsToString[v] = k
	}

	return pr
}

var PermNameValue = map[string]Perm{
	"PermRead":         PermRead,
	"PermOverwrite":    PermOverwrite,
	"PermCreate":       PermCreate,
	"PermDelete":       PermDelete,
	"PermListFiles":    PermListFiles,
	"PermReadMounts":   PermReadMounts,
	"PermUpdateMounts": PermUpdateMounts,
	"PermCreateUser":   PermCreateUser,
	"PermReadUsers":    PermReadUsers,
	"PermWriteUsers":   PermWriteUsers,
	"PermReadPerms":    PermReadPerms,
	"PermWritePerms":   PermWritePerms,
	"PermReadGroups":   PermReadGroups,
	"PermWriteGroups":  PermWriteGroups,
	"PermReadNames":    PermReadNames,
	"PermWriteNames":   PermWriteNames,
	"PermUndefined":    PermUndefined,
	"PermAllDefined":   PermAllDefined,
	"PermAll":          PermAll,
	"PermNone":         PermNone,
	"PermFirst":        PermFirst,
	"PermFiles":        PermFiles,
	"PermWrite":        PermWrite,
	"PermReadOnly":     PermReadOnly,
}

var PermLowerNameValue map[string]Perm

func init() {
	PermLowerNameValue = make(map[string]Perm, len(PermNameValue))

	for k, v := range PermNameValue {
		k = strings.ToLower(k)
		PermLowerNameValue[k] = v
	}
}

var PermShortDefinitions = map[Perm]string{
	PermRead:         "read a file",
	PermOverwrite:    "overwrite an existing file",
	PermCreate:       "create a new file",
	PermDelete:       "delete a file",
	PermListFiles:    "list files",
	PermReadMounts:   "list mounts",
	PermUpdateMounts: "update mounts",
	PermCreateUser:   "create a new user",
	PermReadUsers:    "read all users",
	PermWriteUsers:   "update users",
	PermReadPerms:    "read permissions",
	PermWritePerms:   "write permissions",
	PermReadGroups:   "read all groups",
	PermWriteGroups:  "update groups",
	PermReadNames:    "read all efmrl names",
	PermWriteNames:   "update efmrl names",
	PermUndefined:    "undefined permissions",
	PermAllDefined:   "defined permissions",
	PermAll:          "all permissions",
	PermNone:         "no permissions",
	PermFiles:        "file-oriented permissions",
	PermReadOnly:     "read-only permissions",
}

// PermSimplePerms returns a list of the "simple perms", which are appropriate
// for showing in a UI. It takes out complex ones, e.g. PermFiles, and other
// artifacts that aren't normally used, e.g. PermUndefined.
func PermSimplePerms() []string {
	keys := slices.DeleteFunc(maps.Keys(PermNameValue), func(k string) bool {
		v := PermNameValue[k]
		if v >= PermUndefined {
			return true
		}

		switch k {
		case "PermNone", "PermFirst", "PermAllDefined", "PermFiles",
			"PermWrite", "PermReadOnly":
			return true
		}

		return false
	})

	slices.SortFunc(keys, func(a, b string) int {
		return int(PermNameValue[a] - PermNameValue[b])
	})

	return keys
}

// SimpleNames returns a slice of strings that break down a given perm
func (perm Perm) SimpleNames() []string {
	res := []string{}

	for _, name := range PermSimplePerms() {
		if val := PermNameValue[name]; (val & perm) == val {
			name = strings.TrimPrefix(name, "Perm")
			res = append(res, name)
		}
	}

	return res
}

func SimpleNamesToPerm(names []string) (Perm, error) {
	var res Perm

	for _, name := range names {
		key := cleanup(name)
		perm, found := PermLowerNameValue[key]
		if !found {
			return 0, fmt.Errorf("no such permission %q", name)
		}
		res |= perm
	}

	return res, nil
}

func cleanup(in string) string {
	out := strings.TrimSpace(in)
	out = strings.ToLower(out)
	if len(out) < 4 || out[:4] != "perm" {
		out = "perm" + out
	}

	return out
}
