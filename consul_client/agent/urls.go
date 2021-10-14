package agent

var (
	GET_AGENT_HOSTS       = [3]string{"GET", "/v1/agent/host", "application/json"}
	GET_AGENT_MEMBERS     = [3]string{"GET", "/v1/agent/members", "application/json"}
	GET_AGENT_SELF        = [3]string{"GET", "/v1/agent/self", "application/json"}
	GET_AGENT_RELOAD      = [3]string{"PUT", "/v1/agent/reload", "application/json"}
	GET_AGENT_MAINTENANCE = [3]string{"PUT", "/v1/agent/maintenance", "application/json"}
	GET_AGENT_METRICS     = [3]string{"GET", "/v1/agent/metrics", "application/json"}
	GET_AGENT_MONITOR     = [3]string{"GET", "/v1/agent/monitor", "application/json"}
	GET_AGENT_JOIN        = [3]string{"PUT", "/v1/agent/join", "application/json"}
	GET_AGENT_LEAVE       = [3]string{"PUT", "/v1/agent/leave", "application/json"}
	GET_AGENT_FORCE_LEAVE = [3]string{"PUT", "/v1/agent/force-leave", "application/json"}
)

var (
	UPDATE_TOKEN_DEFAULT      = [3]string{"PUT", "/v1/agent/token/default", "application/json"}
	UPDATE_AGENT_TOKEN        = [3]string{"PUT", "/v1/agent/token/agent", "application/json"}
	UPDATE_AGENT_MASTER_TOKEN = [3]string{"PUT", "/v1/agent/token/agent_master", "application/json"}
	UPDATE_REPLICATION_TOKEN  = [3]string{"PUT", "/v1/agent/token/replication", "application/json"}
)
