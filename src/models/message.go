package models

type Message struct {
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
	Content   string `json:"content"`
	TimeStamp int    `json:"timestamp"`
}
