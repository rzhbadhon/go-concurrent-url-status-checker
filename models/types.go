package models

type RequestBody struct {
	Urls []string `json:"urls"`
}

type Result struct {
	URL    string
	Status string
}