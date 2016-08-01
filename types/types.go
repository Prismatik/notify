package types

type Email struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	From    string `json:"from"`
	To      string `json:"to"`
}

type SMS struct {
	To   string `json:"to"`
	Body string `json:"body"`
	From string `json:"from"`
}
