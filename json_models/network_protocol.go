package json_models

type NetworkProtocol struct {
	Name    string `json:"Name"`
	Enabled bool   `json:"Enabled,string"`
	Port    int    `json:"Port,string"`
}
