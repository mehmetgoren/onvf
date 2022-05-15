package json_models

type NetworkInterface struct {
	Token   string `json:"Token"`
	Enabled bool   `json:"Enabled,string"`
	Info    struct {
		Name      string `json:"Name"`
		HwAddress string `json:"HwAddress"`
		Mtu       int    `json:"MTU,string"`
	} `json:"Info"`
	IPv4 struct {
		Enabled bool `json:"Enabled,string"`
		Config  struct {
			FromDHCP struct {
				Address string `json:"Address"`
			} `json:"FromDHCP"`
			Dhcp bool `json:"DHCP,string"`
		} `json:"Config"`
	} `json:"IPv4"`
}
