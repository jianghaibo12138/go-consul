package connect

import "time"

type CADetail struct {
	ActiveRootID string `json:"ActiveRootID"`
	Roots        []struct {
		ID                  string      `json:"ID"`
		Name                string      `json:"Name"`
		SerialNumber        int         `json:"SerialNumber"`
		SigningKeyID        string      `json:"SigningKeyID"`
		ExternalTrustDomain string      `json:"ExternalTrustDomain"`
		NotBefore           time.Time   `json:"NotBefore"`
		NotAfter            time.Time   `json:"NotAfter"`
		RootCert            string      `json:"RootCert"`
		IntermediateCerts   interface{} `json:"IntermediateCerts"`
		Active              bool        `json:"Active"`
		PrivateKeyType      string      `json:"PrivateKeyType"`
		PrivateKeyBits      int         `json:"PrivateKeyBits"`
		CreateIndex         int         `json:"CreateIndex"`
		ModifyIndex         int         `json:"ModifyIndex"`
	} `json:"Roots"`
}
