package models

import (
	"onvf/json_models"
)

type OnvifModel struct {
	HackResult  *ExecResultView         `json:"hack_result"`
	Onvif       *json_models.TargetInfo `json:"onvif"`
	OnvifParams *OnvifParams            `json:"onvif_params"`
	CreatedAt   string                  `json:"created_at"`
}
