package reps

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"onvf/utils"
	"time"
)

type HeartbeatRepository struct {
	Connection  *redis.Client
	TimeSecond  int64
	ServiceName string
}

func (h *HeartbeatRepository) Start() {
	var dur = time.Duration(h.TimeSecond) * time.Second
	ticker := time.NewTicker(dur)
	//quit := make(chan struct{})
	for {
		select {
		case timeTicker := <-ticker.C:
			heartbeatObj := map[string]interface{}{
				"heartbeat": utils.TimeToString(timeTicker, true),
			}
			h.Connection.HSet(context.Background(), "services:"+h.ServiceName, heartbeatObj)
			log.Println("Heartbeat was beaten at " + timeTicker.Format(time.ANSIC))
			//case <- quit:
			//	ticker.Stop()
			//	return
		}
	}
}
