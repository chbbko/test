package models

type Success struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SuccessArrayResponse struct {
	Success
	Data   interface{} `json:"data"`
	Length int         `json:"length"`
}

type SuccessResponse struct {
	Success
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
