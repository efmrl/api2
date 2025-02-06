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
	Name    string `json:"name"`
	Type    int    `json:"type"`

	Status    string `json:"status"`
	Valid     string `json:"valid,omitempty"`
	Suspended string `json:"suspended,omitempty"`
	Expires   string `json:"expires,omitempty"`

	Usable  bool `json:"usable"`
	Expired bool `json:"expired"`

	Audits []*FundAudit `json:"audits,omitempty"`
}

type FundAudit struct {
	When    string `json:"when"`
	Message string `json:"message"`
}

var FundTypes []string = []string{
	"one shot",
	"subscription",
	"super",
}
