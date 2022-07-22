package JSend

type ErrorResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    *int   `json:"code,omitempty"`
	Data    *T     `json:"data,omitempty"`
}
