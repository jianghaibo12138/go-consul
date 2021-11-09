package consul_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jianghaibo12138/go-consul/consul_client/service"
	"github.com/jianghaibo12138/go-consul/http_client"
	"io/ioutil"
)

func (client *ConsulClient) RegisterService(srv service.RegisterService) (bool, error) {
	url := client.packageRequestTpl(service.SERVICE_REGISTER[1])
	paras := packageQueryBoolParam(false, false, false, false, false, true)
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      service.SERVICE_REGISTER[0],
		Url:         url,
		ContentType: service.SERVICE_REGISTER[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}

	bytes, err := json.Marshal(srv)
	if err != nil {
		return false, err
	}
	response, err := httpClient.Request(bytes)
	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		return false, errors.New(response.Status)
	}
	return true, nil
}

func (client *ConsulClient) DeRegisterService(serviceId string) (bool, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(service.SERVICE_DEREGISTER[1]), serviceId)
	httpClient := http_client.HttpClient{
		Method:      service.SERVICE_DEREGISTER[0],
		Url:         url,
		ContentType: service.SERVICE_DEREGISTER[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}

	bytes, err := json.Marshal([]byte{})
	if err != nil {
		return false, err
	}
	response, err := httpClient.Request(bytes)
	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		return false, errors.New(response.Status)
	}
	return true, nil
}

func (client *ConsulClient) ServiceList() (map[string]interface{}, error) {
	url := client.packageRequestTpl(service.SERVICE_LIST[1])

	httpClient := http_client.HttpClient{
		Method:      service.SERVICE_LIST[0],
		Url:         url,
		ContentType: service.SERVICE_LIST[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var agentService = make(map[string]interface{})
	err = json.Unmarshal(jsonBytes, &agentService)
	if err != nil {
		return nil, err
	}
	return agentService, nil
}

func (client *ConsulClient) GetServiceConfiguration(serviceId, ns string) (*service.Configuration, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(service.SERVICE_CONFIGURATION[1]), serviceId)
	paras := packageQueryStrParam("", "", "", "", "", ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      service.SERVICE_CONFIGURATION[0],
		Url:         url,
		ContentType: service.SERVICE_CONFIGURATION[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var info service.Configuration
	err = json.Unmarshal(jsonBytes, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (client *ConsulClient) GetServiceHealthByName(serviceName, ns string) (*[]service.Health, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(service.SERVICE_HEALTH_BY_NAME[1]), serviceName)
	paras := packageQueryStrParam("", "", "", "", "", ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      service.SERVICE_HEALTH_BY_NAME[0],
		Url:         url,
		ContentType: service.SERVICE_HEALTH_BY_NAME[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var info []service.Health
	err = json.Unmarshal(jsonBytes, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (client *ConsulClient) GetServiceHealthById(serviceId, ns string) (*[]service.Health, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(service.SERVICE_HEALTH_BY_ID[1]), serviceId)
	paras := packageQueryStrParam("", "", "", "", "", ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      service.SERVICE_HEALTH_BY_ID[0],
		Url:         url,
		ContentType: service.SERVICE_HEALTH_BY_ID[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var info []service.Health
	err = json.Unmarshal(jsonBytes, &info)
	if err != nil {
		return nil, err
	}
	return &info, nil
}
