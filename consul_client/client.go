package consul_client

import "fmt"

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

func (client *ConsulClient) packageRequestTpl(urlSuffix string) string {
	var tpl = fmt.Sprintf(HTTP_URL_TPL, client.Host, client.Port, urlSuffix)
	if client.Ssl {
		tpl = fmt.Sprintf(HTTPS_URL_TPL, client.Host, client.Port, urlSuffix)
	}
	return tpl
}
