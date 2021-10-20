package kv

var (
	KV_READ_KEY   = [3]string{"GET", "/v1/kv", "application/json"}
	KV_UPSERT_KEY = [3]string{"PUT", "/v1/kv", "application/json"}
	KV_DELETE_KEY = [3]string{"DELETE", "/v1/kv", "application/json"}
)
