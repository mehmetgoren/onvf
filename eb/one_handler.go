package eb

import (
	"encoding/base64"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log"
	"onvf/models"
	"onvf/reps"
	"onvf/utils"
)

type OneHandler struct {
	Rb        *reps.RepoBucket
	functions map[string]func(*reps.RepoBucket, string) interface{}
}

func (o *OneHandler) Init() {
	o.functions = make(map[string]func(*reps.RepoBucket, string) interface{})
	o.functions["GetTargetInfo"] = getTargetInfo
	o.functions["Reboot"] = reboot
	o.functions["FactoryReset"] = setSystemFactoryDefault
	o.functions["FirmwareUpgrade"] = startFirmwareUpgrade
	o.functions["NetworkDiscovery"] = networkDiscovery
	o.functions["HackTarget"] = hackTarget
}

func (o *OneHandler) Handle(event *redis.Message) error {
	defer utils.HandlePanic()

	model := &models.OnvifEvent{}
	err := json.Unmarshal([]byte(event.Payload), model)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	result := &models.OnvifEvent{}
	result.Type = model.Type
	result.Base64Model = ""
	if fn := o.functions[model.Type]; fn != nil {
		js, err := base64.StdEncoding.DecodeString(model.Base64Model)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		rm := fn(o.Rb, string(js))
		if rm != nil {
			resultJson, err := json.Marshal(&rm)
			if err != nil {
				log.Println(err.Error())
				return err
			}
			result.Base64Model = base64.StdEncoding.EncodeToString(resultJson)
		}
	}
	evt := EventBus{Rb: o.Rb, Channel: "onvif_response"}
	err = evt.Publish(result)
	if err != nil {
		log.Println(err.Error())
	}

	return err
}
