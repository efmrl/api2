package api2

type PostFundReq struct {
	Name         string `json:"name,omitempty"`
	Expires      string `json:"expires,omitempty"`
	Subscription bool   `json:"subscription"`
	Throttled    bool   `json:"throttled"`
	Which        string `json:"which"`
}

type PostFundRes struct {
	ID          string `json:"id"`
	CheckoutURL string `json:"checkoutURL"`
}

type ListFundsRes struct {
	Funds []*Fund `json:"funds"`
}

type GetPortalRes struct {
	URL string `json:"url"`
}

type PostFundReorderReq struct {
	FundID string `json:"fundID"`
	Index  int    `json:"index"`
}

type Fund struct {
	ID      string `json:"id"`
	Created string `json:"created"`
	Name    string `json:"name"`

	Status       string `json:"status"`
	Valid        string `json:"valid,omitempty"`
	Suspended    string `json:"suspended,omitempty"`
	Expires      string `json:"expires,omitempty"`
	Subscription bool   `json:"subscription"`
	Period       string `json:"period"`
	Throttled    bool   `json:"throttled"`

	BurstOnly    bool `json:"burstOnly"`
	PercentAvail int  `json:"percentAvail"`

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
