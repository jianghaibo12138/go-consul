package catalog

type Service struct {
	Node struct {
		ID              string `json:"ID"`
		Node            string `json:"Node"`
		Address         string `json:"Address"`
		Datacenter      string `json:"Datacenter"`
		TaggedAddresses struct {
			Lan string `json:"lan"`
			Wan string `json:"wan"`
		} `json:"TaggedAddresses"`
		Meta struct {
			InstanceType string `json:"instance_type"`
		} `json:"Meta"`
	} `json:"Node"`
	Services []struct {
		ID      string   `json:"ID"`
		Service string   `json:"Service"`
		Tags    []string `json:"Tags"`
		Meta    struct {
			RedisVersion string `json:"redis_version,omitempty"`
		} `json:"Meta"`
		Port            int `json:"Port"`
		TaggedAddresses struct {
			Lan struct {
				Address string `json:"address"`
				Port    int    `json:"port"`
			} `json:"lan"`
			Wan struct {
				Address string `json:"address"`
				Port    int    `json:"port"`
			} `json:"wan"`
		} `json:"TaggedAddresses,omitempty"`
		Namespace string `json:"Namespace,omitempty"`
	} `json:"Services"`
}
