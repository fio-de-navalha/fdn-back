package helpers

type ResponseMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseData struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Error: message,
	}
}
