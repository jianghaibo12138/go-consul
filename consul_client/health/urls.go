package health

var (
	NODE_HEALTH_STATUS    = [3]string{"GET", "/health/node", "application/json"}
	SERVICE_HEALTH_STATUS = [3]string{"GET", "/health/checks", "application/json"}
	SERVICE_INSTANCES     = [3]string{"GET", "/health/service", "application/json"}
)
