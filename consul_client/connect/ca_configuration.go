package connect

type ProviderConfig struct {
	LeafCertTTL         string `json:"LeafCertTTL"`
	PrivateKey          string `json:"PrivateKey"`
	RootCert            string `json:"RootCert"`
	IntermediateCertTTL string `json:"IntermediateCertTTL"`
}

type CAConfiguration struct {
	Provider    string         `json:"Provider"`
	Config      ProviderConfig `json:"Config"`
	CreateIndex int            `json:"CreateIndex"`
	ModifyIndex int            `json:"ModifyIndex"`
}
