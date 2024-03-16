package api2

type PostNameReq struct {
	Name string `json:"name,omitempty"`
}

type PostNameRes struct {
	Name string `json:"name"`
}

type GetNamesRes struct {
	Names []*EfmrlName `json:"names"`
}

type EfmrlName struct {
	Name      string `json:"name"`
	Link      string `json:"link,omitempty"`
	DeleteURL string `json:"deleteURL,omitempty"`
}
