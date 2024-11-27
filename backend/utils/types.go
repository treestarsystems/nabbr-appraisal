package utils

// APIResponse represents a generic API response.
type APIResponse struct {
	Status     string `json:"status"`
	HttpStatus int    `json:"httpStatus"`
	Message    string `json:"message"`
	Payload    any    `json:"payload"`
}

// NewAPIResponse creates a new APIResponse with a dynamic payload type
func NewAPIResponse(status string, httpStatus int, message string, payload any) APIResponse {
	return APIResponse{
		Status:     status,
		HttpStatus: httpStatus,
		Message:    message,
		Payload:    payload,
	}
}
