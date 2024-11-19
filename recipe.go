package api2

type Recipe struct {
	ID      string `json:"id"`
	Created string `json:"created"`
	Source  string `json:"source,omitempty"`

	Name        string `json:"name"`
	ShortDesc   string `json:"short_desc"`
	Description string `json:"description,omitempty"`

	Mounts     map[string]*Mount `json:"mounts,omitempty"`
	Groups     map[string]*Group `json:"groups,omitempty"`
	EfmrlPerms *SpecialPerms     `json:"efmrl_perms,omitempty"`
}

type PostRecipeReq struct {
	Name        string `json:"name"`
	ShortDesc   string `json:"short_desc"`
	Description string `json:"description,omitempty"`
}
