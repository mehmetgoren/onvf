package reps

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"onvf/models"
)

type ConfigRepository struct {
	Connection *redis.Client
}

func (r *ConfigRepository) GetConfig() (*models.Config, error) {
	var config = &models.Config{}
	conn := r.Connection
	key := "config"
	data, err := conn.Get(context.Background(), key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			return config, nil
		} else {
			log.Println("Error getting sources from redis: ", err)
			return nil, err
		}
	}

	err = json.Unmarshal([]byte(data), config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (r *ConfigRepository) SaveConfig(config *models.Config) error {
	conn := r.Connection
	key := "config"
	data, err := json.Marshal(config)
	if err != nil {
		log.Println("Error serializing json: ", err)
		return err
	}

	err = conn.Set(context.Background(), key, data, 0).Err()
	if err != nil {
		log.Println("Error saving ml config to redis: ", err)
		return err
	}

	return nil
}

func (r *ConfigRepository) RestoreConfig() (*models.Config, error) {
	objJson := `{"ai":{"face_recog_mtcnn_threshold":0.86,"face_recog_prob_threshold":0.95,"read_service_overlay":true,"video_clip_duration":10},"device":{"device_name":"gokalp-ubuntu","device_services":[1,2,4,8,16],"device_type":0},"ffmpeg":{"max_operation_retry_count":10000000,"record_concat_limit":1,"record_video_file_indexer_interval":60,"rtmp_server_init_interval":3.0,"start_task_wait_for_interval":1.0,"use_double_quotes_for_path":false,"watch_dog_failed_wait_interval":3.0,"watch_dog_interval":21},"general":{"heartbeat_interval":30,"root_folder_path":"/mnt/sde1"},"jetson":{"model_name":"ssd-mobilenet-v2"},"once_detector":{"imagehash_threshold":3,"psnr_threshold":0.2,"ssim_threshold":0.2},"redis":{"host":"127.0.0.1","port":6379},"source_reader":{"buffer_size":2,"fps":1,"max_retry":150,"max_retry_in":6},"tensorflow":{"cache_folder":"/mnt/sdc1/test_projects/tf_cache","model_name":"efficientdet/lite4/detection"},"torch":{"model_name":"ultralytics/yolov5","model_name_specific":"yolov5x6"}}`

	conn := r.Connection
	key := "config"
	err := conn.Set(context.Background(), key, objJson, 0).Err()
	if err != nil {
		log.Println("Error saving default ml config to redis: ", err)
		return nil, err
	}

	var config = &models.Config{}
	err = json.Unmarshal([]byte(objJson), config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
