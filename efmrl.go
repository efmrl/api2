package api2

type Efmrl struct {
	ID      string `json:"id"`
	Created string `json:"created"`
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
