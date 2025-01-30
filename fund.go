package api2

type PostFundReq struct {
	Type   int    `json:"type"`
	Amount int    `json:"amount"`
	Name   string `json:"name,omitempty"`
}

type PostFundRes struct {
	ID          string `json:"id"`
	CheckoutURL string `json:"checkoutURL"`
}

type ListFundsRes struct {
	Funds []*Fund `json:"funds"`
}

type Fund struct {
	ID      string `json:"id"`
	Created string `json:"created"`
	Name    string `json:"name,omitempty"`
	Type    int    `json:"type"`
	Status  string `json:"status"`
}

var FundTypes []string = []string{
	"one shot",
	"subscription",
	"super",
}
