package model

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ListDataRequestStruct struct {
	Page  int64
	Limit int64
	Order string
}
