package config

type PayloadConfig struct {
	Kind     string `json:"Kind"`
	Name     string `json:"Name"`
	Protocol string `json:"Protocol"`
}

type ResponseConfig struct {
	Kind        string `json:"Kind"`
	Name        string `json:"Name"`
	Protocol    string `json:"Protocol"`
	MeshGateway struct {
	} `json:"MeshGateway"`
	Expose struct {
	} `json:"Expose"`
	CreateIndex int `json:"CreateIndex"`
	ModifyIndex int `json:"ModifyIndex"`
}
