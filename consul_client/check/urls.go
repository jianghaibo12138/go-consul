package check

var (
	CHECK_REGISTER = [3]string{"PUT", "/v1/agent/check/register", "application/json"}
	CHECK_DEREGISTER = [3]string{"PUT", "/v1/agent/check/deregister", "application/json"}
)
