package client_models

type Response struct {
	Statistics map[string]int `json:"statistics"`
	Status     string         `json:"status" enums:"OK"`
}

type BadResponse struct {
	Status     string         `json:"status" example:"NOT OK"`
}
