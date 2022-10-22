package Socket

import (
	"github.com/dunglas/mercure"
	"github.com/oklog/ulid/v2"
	"log"
)

type MercurePusher struct {
	Hub *mercure.Hub
}

func (pusher MercurePusher) Push(topics []string, message string) {
	update := mercure.Update{
		Topics:  topics,
		Private: false,
		Debug:   false,
		Event: mercure.Event{
			Data:  message,
			ID:    ulid.Make().String(),
			Type:  "event",
			Retry: 0,
		},
	}

	// TODO: Determine how to publish update to Mercure hub
	log.Printf("mercure update: %v", update)
}
