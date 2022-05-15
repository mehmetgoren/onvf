package reps

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"onvf/models"
	"onvf/utils"
	"time"
)

type NetworkDiscoveryRepository struct {
	Client *redis.Client
}

func getNetworkKey() string {
	return "onvif:network"
}

func (n *NetworkDiscoveryRepository) Add(results []*models.ExecResultView) (string, error) {
	now := time.Now()
	m := &models.NetworkDiscoveryModel{Results: results, CreatedAt: utils.TimeToString(now, true)}

	jb, err := json.Marshal(m)
	if err != nil {
		return "", err
	}

	js := string(jb)

	var arr interface{} = js

	return n.Client.Set(context.Background(), getNetworkKey(), arr, 0).Result()
}
