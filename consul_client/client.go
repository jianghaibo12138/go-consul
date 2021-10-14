package consul_client

import (
	"amazing2j.com/go-consul/consul_client/agent"
	"amazing2j.com/go-consul/http_client"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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

func (client *ConsulClient) GetHostInfos() (*agent.Agent, error) {
	httpClient := http_client.HttpClient{
		Method:      agent.GET_AGENT_HOSTS[0],
		Url:         fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_HOSTS[1]),
		ContentType: agent.GET_AGENT_HOSTS[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var agentInfo agent.Agent
	err = json.Unmarshal(jsonBytes, &agentInfo)
	if err != nil {
		return nil, err
	}
	return &agentInfo, nil
}

func (client *ConsulClient) GetMembers() (*[]agent.Members, error) {
	httpClient := http_client.HttpClient{
		Method:      agent.GET_AGENT_MEMBERS[0],
		Url:         fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_MEMBERS[1]),
		ContentType: agent.GET_AGENT_MEMBERS[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var members []agent.Members
	err = json.Unmarshal(jsonBytes, &members)
	if err != nil {
		return nil, err
	}
	return &members, nil
}

func (client *ConsulClient) GetSelfConfig() (*agent.SelfConfig, error) {
	httpClient := http_client.HttpClient{
		Method:      agent.GET_AGENT_SELF[0],
		Url:         fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_SELF[1]),
		ContentType: agent.GET_AGENT_SELF[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var configs agent.SelfConfig
	err = json.Unmarshal(jsonBytes, &configs)
	if err != nil {
		return nil, err
	}
	return &configs, nil
}

func (client *ConsulClient) AgentReload() (int, error) {
	httpClient := http_client.HttpClient{
		Method:      agent.GET_AGENT_RELOAD[0],
		Url:         fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_RELOAD[1]),
		ContentType: agent.GET_AGENT_RELOAD[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return 0, err
	}
	return response.StatusCode, nil
}

func (client *ConsulClient) Maintenance(enable bool, reason string) (int, error) {
	httpClient := http_client.HttpClient{
		Method:      agent.GET_AGENT_MAINTENANCE[0],
		Url:         fmt.Sprintf("%s?enable=%v&reason=%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_MAINTENANCE[1]), enable, reason),
		ContentType: agent.GET_AGENT_MAINTENANCE[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return 0, err
	}
	return response.StatusCode, nil
}

func (client *ConsulClient) Metrics(format string) (*agent.Metrics, error) {
	url := fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_METRICS[1])

	if len(format) != 0 {
		url = fmt.Sprintf("%s?format=%s", url, format)
	}

	httpClient := http_client.HttpClient{
		Method:      agent.GET_AGENT_METRICS[0],
		Url:         url,
		ContentType: agent.GET_AGENT_METRICS[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var metrics agent.Metrics
	err = json.Unmarshal(jsonBytes, &metrics)
	if err != nil {
		return nil, err
	}
	return &metrics, nil
}

func (client *ConsulClient) MetricsWithFormat(format string) (*[]byte, error) {
	url := fmt.Sprintf("%s?format=%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_METRICS[1]), format)
	httpClient := http_client.HttpClient{
		Method:      agent.GET_AGENT_METRICS[0],
		Url:         url,
		ContentType: "text/plain; version=0.0.4; charset=utf-8",
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &bytes, nil
}

func (client *ConsulClient) Monitor(logJson bool, logLevel string, logChannel chan []byte) error {
	// TODO realize
	// url := fmt.Sprintf("%s?logjson=%v&loglevel=%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_MONITOR[1]), logJson, logLevel)
	//
	// httpClient := http_client.HttpClient{
	//	Method:      agent.GET_AGENT_MONITOR[0],
	//	Url:         url,
	//	ContentType: agent.GET_AGENT_MONITOR[2],
	//	Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	// }
	// response, err := httpClient.Request([]byte{})
	// if err != nil {
	//	return err
	// }
	// for {
	//	jsonBytes, err := ioutil.ReadAll(response.Body)
	//	if err != nil {
	//		return err
	//	}
	//	logChannel <- jsonBytes
	// }
	return nil
}

func (client *ConsulClient) Join(address string, wan bool) (*[]byte, error) {
	url := fmt.Sprintf("%s/%s/%+v", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_JOIN[1]), address, wan)
	httpClient := http_client.HttpClient{
		Method:      agent.GET_AGENT_JOIN[0],
		Url:         url,
		ContentType: agent.GET_AGENT_JOIN[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &bytes, nil
}

func (client *ConsulClient) Leave() (*[]byte, error) {
	httpClient := http_client.HttpClient{
		Method:      agent.GET_AGENT_LEAVE[0],
		Url:         fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_LEAVE[1]),
		ContentType: agent.GET_AGENT_LEAVE[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &bytes, nil
}

func (client *ConsulClient) ForceLeave(node string, prune bool) (*[]byte, error) {
	url := fmt.Sprintf("%s/%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_FORCE_LEAVE[1]), node)
	if prune {
		url = fmt.Sprintf("%s?prnue", url)
	}
	httpClient := http_client.HttpClient{
		Method:      agent.GET_AGENT_FORCE_LEAVE[0],
		Url:         url,
		ContentType: agent.GET_AGENT_FORCE_LEAVE[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &bytes, nil
}

// TODO Add token update
