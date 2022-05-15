package eb

import "github.com/go-redis/redis/v8"

type EventHandler interface {
	Handle(event *redis.Message) error
}
