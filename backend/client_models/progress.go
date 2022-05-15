package client_models

type Response struct {
	Statistics map[string]int `json:"statistics"`
	Status     string         `json:"status"`
}