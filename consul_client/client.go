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
		Url:         fmt.Sprintf("%s?token=%v&reason=%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, agent.GET_AGENT_MAINTENANCE[1]), enable, reason),
		ContentType: agent.GET_AGENT_MAINTENANCE[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return 0, err
	}
	return response.StatusCode, nil
}
