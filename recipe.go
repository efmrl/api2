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

	ProtoEfmrl  string `json:"protoEfmrl,omitempty"`
	OwningEfmrl string `json:"owningEfmrl,omitempty"`
}

type PostRecipeReq struct {
	Name        string `json:"name"`
	ShortDesc   string `json:"short_desc"`
	Description string `json:"description,omitempty"`
}

type GetRecipesRes struct {
	OwnedRecipies []string `json:"ownedRecipes,omitempty"`
	ProtoRecipes  []string `json:"protoRecipes,omitempty"`
}
