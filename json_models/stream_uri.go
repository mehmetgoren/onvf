package json_models

type StreamUri struct {
	URI                 string `json:"URI"`
	InvalidAfterConnect bool   `json:"InvalidAfterConnect,string"`
	InvalidAfterReboot  bool   `json:"InvalidAfterReboot,string"`
	Timeout             string `json:"Timeout"`
}
