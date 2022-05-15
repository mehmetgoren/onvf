package reps

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"onvf/json_models"
	"onvf/models"
	"onvf/utils"
	"time"
)

type OnvifRepository struct {
	Client *redis.Client
}

func getTargetKey(address string) string {
	ip := utils.ParseIp(address)
	return "onvif:" + ip
}

func (o *OnvifRepository) Get(address string) (*models.OnvifModel, error) {
	conn := o.Client
	key := getTargetKey(address)
	js, err := conn.Get(context.Background(), key).Result()
	if len(js) == 0 {
		return nil, nil
	}
	model := &models.OnvifModel{}
	err = json.Unmarshal([]byte(js), model)
	if err != nil {
		log.Println("Error getting target model from redis: ", err)
		return nil, err
	}

	return model, nil
}

func (o *OnvifRepository) saveTargetModel(params *models.OnvifParams, er *models.ExecResultView, onvif *json_models.TargetInfo) (string, error) {
	address := params.Address
	m, _ := o.Get(address)
	if m == nil {
		m = &models.OnvifModel{}
		now := time.Now()
		m.CreatedAt = utils.TimeToString(now, true)
	}
	m.OnvifParams = params
	if er != nil {
		m.HackResult = er
	}
	if onvif != nil {
		m.Onvif = onvif
	}

	key := getTargetKey(address)
	jb, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	js := string(jb)

	return o.Client.Set(context.Background(), key, js, 0).Result()
}

func (o *OnvifRepository) AddWithExecResult(params *models.OnvifParams, er *models.ExecResultView) (string, error) {
	return o.saveTargetModel(params, er, nil)
}

func (o *OnvifRepository) AddWithOnvif(params *models.OnvifParams, onvif *json_models.TargetInfo) (string, error) {
	return o.saveTargetModel(params, nil, onvif)
}
