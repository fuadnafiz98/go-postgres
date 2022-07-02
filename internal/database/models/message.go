package models

type Message struct {
	ID   uint   `json:"id" grom:"primary_key"`
	Text string `json:"text"`
}
