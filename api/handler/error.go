package handler

type ErrorRes struct {
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
}
