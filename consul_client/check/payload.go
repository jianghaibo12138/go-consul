package check

type RegisterPayload struct {
	ID                             string   `json:"ID"`
	Name                           string   `json:"Name"`
	Notes                          string   `json:"Notes"`
	DeregisterCriticalServiceAfter string   `json:"DeregisterCriticalServiceAfter"`
	Args                           []string `json:"Args"`
	DockerContainerID              string   `json:"DockerContainerID"`
	Shell                          string   `json:"Shell"`
	HTTP                           string   `json:"HTTP"`
	Method                         string   `json:"Method"`
	Header                         struct {
		ContentType []string `json:"Content-Type"`
	} `json:"Header"`
	Body          string `json:"Body"`
	TCP           string `json:"TCP"`
	Interval      string `json:"Interval"`
	Timeout       string `json:"Timeout"`
	TLSSkipVerify bool   `json:"TLSSkipVerify"`
}
