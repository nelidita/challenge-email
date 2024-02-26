package models

type Email struct {
	Id       string `json:"id"`
	From     string `json:"from"`
	To       string `json:"to"`
	Content  string `json:"content"`
	Subject  string `json:"subject"`
	Date     string `json:"date"`
	Filepath string `json:"filepath"`
}

type CreateIndexRequest struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}
