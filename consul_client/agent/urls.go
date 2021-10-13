package agent

var (
	GET_AGENT_HOSTS       = [3]string{"GET", "/v1/agent/host", "application/json"}
	GET_AGENT_MEMBERS     = [3]string{"GET", "/v1/agent/members", "application/json"}
	GET_AGENT_SELF        = [3]string{"GET", "/v1/agent/self", "application/json"}
	GET_AGENT_RELOAD      = [3]string{"PUT", "/v1/agent/reload", "application/json"}
	GET_AGENT_MAINTENANCE = [3]string{"PUT", "/v1/agent/maintenance", "application/json"}
	GET_AGENT_METRICS     = [3]string{"GET", "/v1/agent/metrics", "application/json"}
	GET_AGENT_MONITOR     = [3]string{"GET", "/v1/agent/monitor", "application/json"}
	GET_AGENT_JOIN     = [3]string{"PUT", "/v1/agent/join", "application/json"}
)
