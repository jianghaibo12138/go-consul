package catalog

var (
	CATALOG_REGISTER      = [3]string{"PUT", "/v1/catalog/register", "application/json"}
	CATALOG_DEREGISTER    = [3]string{"PUT", "/v1/catalog/deregister", "application/json"}
	CATALOG_DATACENTERS   = [3]string{"GET", "/v1/catalog/datacenters", "application/json"}
	CATALOG_NODES         = [3]string{"GET", "/v1/catalog/nodes", "application/json"}
	CATALOG_SERVICES      = [3]string{"GET", "/v1/catalog/services", "application/json"}
	CATALOG_SERVICE_NODES = [3]string{"GET", "/v1/catalog/service", "application/json"}
	CATALOG_CONNECT_NODES = [3]string{"GET", "/v1/catalog/connect", "application/json"}
	CATALOG_NODE_NODES    = [3]string{"GET", "/v1/catalog/node", "application/json"}
	CATALOG_NODE_SERVICES    = [3]string{"GET", "/v1/catalog/node-services", "application/json"}
)
