package service

var (
	SERVICE_LIST           = [3]string{"GET", "/v1/agent/services", "application/json"}
	SERVICE_CONFIGURATION  = [3]string{"GET", "/v1/agent/service", "application/json"}
	SERVICE_HEALTH_BY_NAME = [3]string{"GET", "/v1/agent/health/service/name", "application/json"}
	SERVICE_HEALTH_BY_ID   = [3]string{"GET", "/v1/agent/health/service/id", "application/json"}
	SERVICE_REGISTER       = [3]string{"PUT", "/v1/agent/service/register", "application/json"}
	SERVICE_DEREGISTER     = [3]string{"PUT", "/v1/agent/service/deregister", "application/json"}
)
