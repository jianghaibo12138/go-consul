package catalog

type Node struct {
	ID              string `json:"ID"`
	Node            string `json:"Node"`
	Address         string `json:"Address"`
	Datacenter      string `json:"Datacenter"`
	TaggedAddresses struct {
		Lan     string `json:"lan"`
		LanIpv4 string `json:"lan_ipv4"`
		Wan     string `json:"wan"`
		WanIpv4 string `json:"wan_ipv4"`
	} `json:"TaggedAddresses"`
	Meta struct {
		ConsulNetworkSegment string `json:"consul-network-segment"`
	} `json:"Meta"`
	CreateIndex int `json:"CreateIndex"`
	ModifyIndex int `json:"ModifyIndex"`
}
