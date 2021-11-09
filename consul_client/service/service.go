package service

type RegisterService struct {
	ID      string   `json:"ID"`
	Name    string   `json:"Name"`
	Tags    []string `json:"Tags"`
	Address string   `json:"Address"`
	Port    int      `json:"Port"`
	Meta    struct {
		RedisVersion string `json:"redis_version"`
	} `json:"Meta"`
	EnableTagOverride bool `json:"EnableTagOverride"`
	Check             struct {
		DeregisterCriticalServiceAfter string   `json:"DeregisterCriticalServiceAfter"`
		Args                           []string `json:"Args"`
		Interval                       string   `json:"Interval"`
		Timeout                        string   `json:"Timeout"`
	} `json:"Check"`
	Weights struct {
		Passing int `json:"Passing"`
		Warning int `json:"Warning"`
	} `json:"Weights"`
}

type Configuration struct {
	Kind            string      `json:"Kind"`
	ID              string      `json:"ID"`
	Service         string      `json:"Service"`
	Tags            interface{} `json:"Tags"`
	Meta            interface{} `json:"Meta"`
	Namespace       string      `json:"Namespace"`
	Port            int         `json:"Port"`
	Address         string      `json:"Address"`
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
	Weights struct {
		Passing int `json:"Passing"`
		Warning int `json:"Warning"`
	} `json:"Weights"`
	EnableTagOverride bool   `json:"EnableTagOverride"`
	Datacenter        string `json:"Datacenter"`
	ContentHash       string `json:"ContentHash"`
	Proxy             struct {
		DestinationServiceName string `json:"DestinationServiceName"`
		DestinationServiceID   string `json:"DestinationServiceID"`
		LocalServiceAddress    string `json:"LocalServiceAddress"`
		LocalServicePort       int    `json:"LocalServicePort"`
		Mode                   string `json:"Mode"`
		TransparentProxy       struct {
			OutboundListenerPort int `json:"OutboundListenerPort"`
		} `json:"TransparentProxy"`
		Config struct {
			Foo string `json:"foo"`
		} `json:"Config"`
		Upstreams []struct {
			DestinationType string `json:"DestinationType"`
			DestinationName string `json:"DestinationName"`
			LocalBindPort   int    `json:"LocalBindPort"`
		} `json:"Upstreams"`
	} `json:"Proxy"`
}

type Health struct {
	AggregatedStatus string `json:"AggregatedStatus"`
	Service          struct {
		ID      string   `json:"ID"`
		Service string   `json:"Service"`
		Tags    []string `json:"Tags"`
		Meta    struct {
		} `json:"Meta"`
		Port            int    `json:"Port"`
		Address         string `json:"Address"`
		SocketPath      string `json:"SocketPath"`
		TaggedAddresses struct {
			Lan struct {
				Address string `json:"Address"`
				Port    int    `json:"Port"`
			} `json:"lan"`
			Wan struct {
				Address string `json:"Address"`
				Port    int    `json:"Port"`
			} `json:"wan"`
		} `json:"TaggedAddresses"`
		Weights struct {
			Passing int `json:"Passing"`
			Warning int `json:"Warning"`
		} `json:"Weights"`
		EnableTagOverride bool   `json:"EnableTagOverride"`
		Namespace         string `json:"Namespace"`
		Datacenter        string `json:"Datacenter"`
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
		Type        string   `json:"Type"`
		Namespace   string   `json:"Namespace"`
		Definition  struct {
			Interval                       string      `json:"Interval"`
			Timeout                        string      `json:"Timeout"`
			DeregisterCriticalServiceAfter string      `json:"DeregisterCriticalServiceAfter"`
			HTTP                           string      `json:"HTTP"`
			Header                         interface{} `json:"Header"`
			Method                         string      `json:"Method"`
			Body                           string      `json:"Body"`
			TLSServerName                  string      `json:"TLSServerName"`
			TLSSkipVerify                  bool        `json:"TLSSkipVerify"`
			TCP                            string      `json:"TCP"`
		} `json:"Definition"`
		CreateIndex int `json:"CreateIndex"`
		ModifyIndex int `json:"ModifyIndex"`
	} `json:"Checks"`
}
