package reps

import (
	"github.com/go-redis/redis/v8"
)

type RepoBucket struct {
	connMain         *redis.Client
	PubSubConnection *redis.Client
	ConfigRep        *ConfigRepository
}

func (r *RepoBucket) Init() *RepoBucket {
	r.connMain = createRedisConnection(MAIN)
	r.PubSubConnection = createRedisConnection(EVENTBUS)
	r.ConfigRep = &ConfigRepository{Connection: r.connMain}

	return r
}

func (r *RepoBucket) GetMainConnection() *redis.Client {
	return r.connMain
}
