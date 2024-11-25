package utils

// APIResponse represents a generic API response.
type APIResponse struct {
	Status     string        `json:"status"`
	HttpStatus int           `json:"httpStatus"`
	Message    string        `json:"message"`
	Payload    []interface{} `json:"payload"`
}
