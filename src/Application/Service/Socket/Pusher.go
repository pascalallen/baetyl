package Socket

type Pusher interface {
	Push(topics []string, message string)
}
