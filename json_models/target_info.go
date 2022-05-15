package json_models

type TargetInfo struct {
	DeviceInfo *DeviceInfo `json:"device_info"`

	IsDiscoverable bool   `json:"is_discoverable"`
	HostName       string `json:"host_name"`

	IPAddresses []string `json:"ip_addresses"`

	//GetNetworkInterfaces
	HwAddress string `json:"hw_address"`

	//GetNetworkProtocols
	HttpPort  int `json:"http_port"`
	HttpsPort int `json:"https_port"`
	RtspPort  int `json:"rtsp_port"`

	//GetSystemDateAndTime
	LocalDatetime string `json:"local_datetime"`

	//GetSystemLog
	Logs string `json:"logs"`

	Users []*User `json:"users"`

	StreamUri string `json:"stream_uri"`
}
