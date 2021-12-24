package models

type Result struct {
	Result string `json:"result"`
}

type Name struct {
	OldNameEvent string `json:"old_name_event"`
}

type Event struct {
	UserID     int    `json:"user_id"`
	NameEvent  string `json:"name_event"`
	Date       string `json:"date"`
	Descripton string `json:"descripton"`
}
