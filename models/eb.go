package models

import "encoding/json"

type OnvifParams struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Asset struct {
	Address string `json:"address"`
}

type OnvifEvent struct {
	Type        string `json:"type"`
	Base64Model string `json:"base64_model"`
}

func (o OnvifEvent) MarshalBinary() ([]byte, error) {
	return json.Marshal(o)
}
