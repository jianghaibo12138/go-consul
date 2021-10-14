package catalog

type RegisterCatalog struct {
	Datacenter      string `json:"Datacenter"`
	ID              string `json:"ID"`
	Node            string `json:"Node"`
	Address         string `json:"Address"`
	TaggedAddresses struct {
		Lan string `json:"lan"`
		Wan string `json:"wan"`
	} `json:"TaggedAddresses"`
	NodeMeta struct {
		Somekey string `json:"somekey"`
	} `json:"NodeMeta"`
	Service struct {
		ID              string   `json:"ID"`
		Service         string   `json:"Service"`
		Tags            []string `json:"Tags"`
		Address         string   `json:"Address"`
		TaggedAddresses struct {
			Lan struct {
				Address string `json:"address"`
				Port    int    `json:"port"`
			} `json:"lan"`
			Wan struct {
				Address string `json:"address"`
				Port    int    `json:"port"`
			} `json:"wan"`
		} `json:"TaggedAddresses"`
		Meta struct {
			RedisVersion string `json:"redis_version"`
		} `json:"Meta"`
		Port      int    `json:"Port"`
		Namespace string `json:"Namespace"`
	} `json:"Service"`
	Check struct {
		Node       string `json:"Node"`
		CheckID    string `json:"CheckID"`
		Name       string `json:"Name"`
		Notes      string `json:"Notes"`
		Status     string `json:"Status"`
		ServiceID  string `json:"ServiceID"`
		Definition struct {
			TCP                            string `json:"TCP"`
			Interval                       string `json:"Interval"`
			Timeout                        string `json:"Timeout"`
			DeregisterCriticalServiceAfter string `json:"DeregisterCriticalServiceAfter"`
		} `json:"Definition"`
		Namespace string `json:"Namespace"`
	} `json:"Check"`
	SkipNodeUpdate bool `json:"SkipNodeUpdate"`
}

type DeRegisterCatalog struct {
	Datacenter string `json:"Datacenter"`
	Node       string `json:"Node"`
	ServiceID  string `json:"ServiceID"`
	Namespace  string `json:"Namespace"`
}

type NodeServices struct {
	Node     Node      `json:"Node"`
	Services []Service `json:"Services"`
}
