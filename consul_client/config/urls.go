package config

var (
	CONFIG_UPSERT      = [3]string{"PUT", "/v1/config", "application/json"}
	CONFIG_GETTER      = [3]string{"GET", "/v1/config", "application/json"}
	CONFIG_LIST_GETTER = [3]string{"GET", "/v1/config", "application/json"}
	CONFIG_DELETE      = [3]string{"DELETE", "/v1/config", "application/json"}
)
