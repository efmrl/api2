package api2

type PostFundReq struct {
	Type int    `json:"type"`
	Name string `json:"name,omitempty"`
}

type PostFundRes struct {
	ID          string `json:"id"`
	CheckoutURL string `json:"jsonURL"`
}

type Fund struct {
	ID      string `json:"id"`
	Created string `json:"created"`
	Name    string `json:"name,omitempty"`
	Type    int    `json:"type"`
}

var FundTypes []string = []string{
	"one shot",
	"subscription",
	"super",
}
