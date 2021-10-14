package consul_client

type ConsulClientOption struct {
	Version string `json:"version"`
}

const (
	CONSUL_TOKEN_KEY     = "X-Consul-Token"
	CONSUL_NAMESPACE_KEY = "X-Consul-Namespace"
)
