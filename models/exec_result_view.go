package models

type ExecResultView struct {
	Device   string   `json:"device"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Routes   []string `json:"route"`
	Address  string   `json:"address"`
	Port     uint16   `json:"port"`

	CredentialsFound bool `json:"credentials_found"`
	RouteFound       bool `json:"route_found"`
	Available        bool `json:"available"`

	AuthenticationType string `json:"authentication_type"`
}
