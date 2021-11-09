package common

type LiveCheck struct {
	DeregisterCriticalServiceAfter string      `json:"DeregisterCriticalServiceAfter"`
	Args                           []string    `json:"Args"`
	Interval                       string      `json:"Interval"`
	Timeout                        string      `json:"Timeout"`
	HTTP                           string      `json:"HTTP"`
	Header                         interface{} `json:"Header"`
	Method                         string      `json:"Method"`
	Body                           string      `json:"Body"`
	TLSServerName                  string      `json:"TLSServerName"`
	TLSSkipVerify                  bool        `json:"TLSSkipVerify"`
	TCP                            string      `json:"TCP"`
}

type Weights struct {
	Passing int `json:"Passing"`
	Warning int `json:"Warning"`
}

type Addr struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type TaggedAddresses struct {
	Lan Addr `json:"lan"`
	Wan Addr `json:"wan"`
}

type TransparentProxy struct {
	OutboundListenerPort int `json:"OutboundListenerPort"`
}

type Upstreams struct {
	DestinationType string `json:"DestinationType"`
	DestinationName string `json:"DestinationName"`
	LocalBindPort   int    `json:"LocalBindPort"`
}

type Proxy struct {
	DestinationServiceName string                 `json:"DestinationServiceName"`
	DestinationServiceID   string                 `json:"DestinationServiceID"`
	LocalServiceAddress    string                 `json:"LocalServiceAddress"`
	LocalServicePort       int                    `json:"LocalServicePort"`
	Mode                   string                 `json:"Mode"`
	TransparentProxy       TransparentProxy       `json:"TransparentProxy"`
	Config                 map[string]interface{} `json:"Config"`
	Upstreams              Upstreams              `json:"Upstreams"`
}

type CheckDetail struct {
	Node        string           `json:"Node"`
	CheckID     string           `json:"CheckID"`
	Name        string           `json:"Name"`
	Status      string           `json:"Status"`
	Notes       string           `json:"Notes"`
	Output      string           `json:"Output"`
	ServiceID   string           `json:"ServiceID"`
	ServiceName string           `json:"ServiceName"`
	ServiceTags []string         `json:"ServiceTags"`
	Type        string           `json:"Type"`
	Namespace   string           `json:"Namespace"`
	Definition  LiveCheck `json:"Definition"`
	CreateIndex int              `json:"CreateIndex"`
	ModifyIndex int              `json:"ModifyIndex"`
}
