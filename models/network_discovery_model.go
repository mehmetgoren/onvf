package models

type NetworkDiscoveryModel struct {
	Results   []*ExecResultView `json:"results"`
	CreatedAt string            `json:"created_at"`
}
