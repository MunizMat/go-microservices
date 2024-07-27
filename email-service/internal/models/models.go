package models

type Email struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Content string `json:"content"`
	UserId  string `json:"userId"`
	Id      string `json:"id"`
}
