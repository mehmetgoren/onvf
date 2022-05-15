package eb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"onvf/reps"
	"onvf/utils"
)

type EventBus struct {
	Rb      *reps.RepoBucket `json:"-"`
	Channel string
}

func (eb *EventBus) Publish(event interface{}) error {
	return eb.Rb.PubSubConnection.Publish(context.Background(), eb.Channel, event).Err()
}

func (eb *EventBus) Subscribe(handler EventHandler) error {
	utils.HandlePanic()

	pong, err := eb.Rb.PubSubConnection.Ping(context.Background()).Result()
	if err != nil {
		log.Println("ping has been failed, exiting now...")
		panic(err)
		return err
	}

	log.Println("ping: " + pong + " for " + eb.Channel)
	log.Println("redis pubsub is listening for " + eb.Channel)

	channel := eb.Channel
	subscribe := eb.Rb.PubSubConnection.Subscribe(context.Background(), channel)
	subscriptions := subscribe.ChannelWithSubscriptions(context.Background(), 1)
	for {
		select {
		case sub := <-subscriptions:
			var message, isRedisMessage = sub.(*redis.Message)
			if !isRedisMessage {
				continue
			}
			go func() {
				err := handler.Handle(message)
				if err != nil {
					log.Println(err.Error())
				}
			}()
		}
	}
	return nil
}
