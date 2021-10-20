package catalog

type Gateway struct {
	Gateway struct {
		Name      string `json:"Name"`
		Namespace string `json:"Namespace"`
	} `json:"Gateway"`
	Service struct {
		Name      string `json:"Name"`
		Namespace string `json:"Namespace"`
	} `json:"Service"`
	GatewayKind string `json:"GatewayKind"`
	CAFile      string `json:"CAFile,omitempty"`
	CertFile    string `json:"CertFile,omitempty"`
	KeyFile     string `json:"KeyFile,omitempty"`
	SNI         string `json:"SNI,omitempty"`
	CreateIndex int    `json:"CreateIndex"`
	ModifyIndex int    `json:"ModifyIndex"`
}
