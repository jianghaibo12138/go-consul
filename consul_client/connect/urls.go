package connect

var (
	GET_CONNECT_CA_ROOT             = [3]string{"GET", "/v1/connect/ca/roots", "application/json"}
	GET_CONNECT_CA_CONFIGURATION    = [3]string{"GET", "/v1/connect/ca/configuration", "application/json"}
	UPDATE_CONNECT_CA_CONFIGURATION = [3]string{"PUT", "/v1/connect/ca/configuration", "application/json"}
	UPSERT_INTENTIONS_BY_NAME       = [3]string{"PUT", "/v1/connect/intentions/exact", "application/json"}
	CREATE_INTENTIONS_WITH_ID       = [3]string{"POST", "/v1/connect/intentions", "application/json"}
	UPDATE_INTENTIONS_BY_ID         = [3]string{"PUT", "/v1/connect/intentions", "application/json"}
	GET_INTENTION_BY_NAME           = [3]string{"GET", "/v1/connect/intentions/exact", "application/json"}
	GET_INTENTION_BY_ID             = [3]string{"GET", "/v1/connect/intentions", "application/json"}
	GET_INTENTION_LIST              = [3]string{"GET", "/v1/connect/intentions", "application/json"}
	DELETE_INTENTION_BY_NAME        = [3]string{"DELETE", "/v1/connect/intentions/exact", "application/json"}
	DELETE_INTENTION_BY_ID          = [3]string{"DELETE", "/v1/connect/intentions", "application/json"}
	CHECK_INTENTION_RESULT          = [3]string{"GET", "/v1/connect/intentions/check", "application/json"}
)
