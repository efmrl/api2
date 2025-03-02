package api2

type GetPasskeysRes struct {
	Passkeys []*Passkey `json:"passkeys"`
}

type Passkey struct {
	ID       string `json:"id"`
	Created  string `json:"created"`
	LastUsed string `json:"lastUsed,omitempty"`
}
