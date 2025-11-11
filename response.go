package main

type SuccessResponse struct {
	Words   []string `json:"words"`
	Results int      `json:"results"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
