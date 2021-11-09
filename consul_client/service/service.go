package service

import "github.com/jianghaibo12138/go-consul/consul_client/common"

type Payload struct {
	ID                string                 `json:"ID"`
	Name              string                 `json:"Name"`
	Tags              []string               `json:"Tags"`
	Address           string                 `json:"Address"`
	Port              int                    `json:"Port"`
	Meta              map[string]interface{} `json:"Meta"`
	EnableTagOverride bool                   `json:"EnableTagOverride"`
	Check             common.LiveCheck       `json:"Check"`
	Weights           common.Weights         `json:"Weights"`
}

type Detail struct {
	ID                string                 `json:"ID"`
	Service           string                 `json:"Service"`
	Tags              []interface{}          `json:"Tags"`
	TaggedAddresses   common.TaggedAddresses `json:"TaggedAddresses"`
	Meta              map[string]interface{} `json:"Meta"`
	Namespace         string                 `json:"Namespace"`
	Port              int                    `json:"Port"`
	Address           string                 `json:"Address"`
	EnableTagOverride bool                   `json:"EnableTagOverride"`
	Datacenter        string                 `json:"Datacenter"`
	Weights           common.Weights         `json:"Weights"`
}

type Configuration struct {
	Kind              string                 `json:"Kind"`
	ID                string                 `json:"ID"`
	Service           string                 `json:"Service"`
	Tags              interface{}            `json:"Tags"`
	Meta              interface{}            `json:"Meta"`
	Namespace         string                 `json:"Namespace"`
	Port              int                    `json:"Port"`
	Address           string                 `json:"Address"`
	TaggedAddresses   common.TaggedAddresses `json:"TaggedAddresses"`
	Weights           common.Weights         `json:"Weights"`
	EnableTagOverride bool                   `json:"EnableTagOverride"`
	Datacenter        string                 `json:"Datacenter"`
	ContentHash       string                 `json:"ContentHash"`
	Proxy             common.Proxy           `json:"Proxy"`
}

type Health struct {
	AggregatedStatus string               `json:"AggregatedStatus"`
	Service          Detail               `json:"Service"`
	Checks           []common.CheckDetail `json:"Checks"`
}
