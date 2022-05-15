package json_models

type Service struct {
	Namespace string `json:"Namespace"`
	XAddr     string `json:"XAddr"`
	Version   struct {
		Major string `json:"Major"`
		Minor string `json:"Minor"`
	} `json:"Version"`
}
