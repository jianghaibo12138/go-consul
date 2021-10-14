package consul_client

type ConsulClient struct {
	Host  string `json:"host"`
	Port  int    `json:"port"`
	Token string `json:"token"`
	Ssl   bool   `json:"ssl"`
}

const (
	HTTP_URL_TPL  = "http://%s:%d%s"
	HTTPS_URL_TPL = "http://%s:%d%s"
)

func (client *ConsulClient) packageRequestTpl() string {
	var tpl = HTTP_URL_TPL
	if client.Ssl {
		tpl = HTTPS_URL_TPL
	}
	return tpl
}
