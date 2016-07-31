package types

type Email struct {
	Subject string
	Body    string
	From    string
	To      string
}

type SMS struct {
	To   string
	Body string
	From string
}
