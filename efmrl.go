package api2

type Efmrl struct {
	ID      string `json:"id"`
	Created string `json:"created"`
}

type GetEfmrlMDRes struct {
	CanonicalURL string `json:"canonicalURL"`
	APIPrefix    string `json:"apiPrefix"`
}

type EfmrlSpecialPaths struct {
	Wildcard string `json:"wildcard"`
}

type GetCreationSessionReq struct {
	Email string `json:"email"`
}

type ReserveNameReq struct {
	Name string `json:"name"`
}

type ReserveNameRes struct {
	Expires string `json:"expires"`
}

type FindEfmrlsReq struct {
	EmailAddr string `json:"emailAddr,omitempty"`

	MaxCount   int32  `json:"maxCount,omitempty"`
	StartAfter string `json:"startAfter,omitempty"`
}

type FindEfmrlsRes struct {
	Efmrls []*struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"efmrls"`

	ExclNext string `json:"exclNext,omitempty"`
	Count    int32  `json:"count"`
}

type PostEfmrlReq struct {
	OwnerName string       `json:"ownerName,omitempty"`
	Name      string       `json:"name,omitempty"`
	Recipe    string       `json:"recipe,omitempty"`
	Funding   *PostFundReq `json:"funding,omitempty"`
}

type PostEfmrlRes struct {
	EfmrlID     string `json:"efmrlID"`
	FundID      string `json:"fundID"`
	CheckoutURL string `json:"checkoutURL"`
}
