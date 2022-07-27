package JSend

type FailResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}
