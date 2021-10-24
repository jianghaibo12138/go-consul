package health

type ServiceInstance struct {
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
		Port    int `json:"Port"`
		Weights struct {
			Passing int `json:"Passing"`
			Warning int `json:"Warning"`
		} `json:"Weights"`
		Namespace string `json:"Namespace"`
	} `json:"Service"`
	Checks []struct {
		Node        string   `json:"Node"`
		CheckID     string   `json:"CheckID"`
		Name        string   `json:"Name"`
		Status      string   `json:"Status"`
		Notes       string   `json:"Notes"`
		Output      string   `json:"Output"`
		ServiceID   string   `json:"ServiceID"`
		ServiceName string   `json:"ServiceName"`
		ServiceTags []string `json:"ServiceTags"`
		Namespace   string   `json:"Namespace"`
	} `json:"Checks"`
}