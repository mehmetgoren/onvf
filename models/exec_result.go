package models

const (
	authNone   = 0
	authBasic  = 1
	authDigest = 2
)

// ExecResult represents a camera's RTSP stream
type ExecResult struct {
	Device   string   `json:"device"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Routes   []string `json:"route"`
	Address  string   `json:"address"`
	Port     uint16   `json:"port"`

	CredentialsFound bool `json:"credentials_found"`
	RouteFound       bool `json:"route_found"`
	Available        bool `json:"available"`

	AuthenticationType int `json:"authentication_type"`
}

func (e *ExecResult) GetAuthenticationTypeString() string {
	switch e.AuthenticationType {
	case authNone:
		return "no"
	case authBasic:
		return "basic"
	case authDigest:
		return "digest"
	default:
		return ""
	}
}
