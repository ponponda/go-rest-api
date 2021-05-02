package models

// swagger:parameters ApiResponse
type ApiResponse struct {
	StatusCode int
	Message    string
	Payload    interface{}
}
