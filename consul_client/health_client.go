package consul_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jianghaibo12138/go-consul/consul_client/health"
	"github.com/jianghaibo12138/go-consul/http_client"
	"io/ioutil"
)

func (client *ConsulClient) NodeHealthStatus(node, ns, dc, filter string) ([]health.NodeStatus, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(health.NODE_HEALTH_STATUS[1]), node)
	paras := packageQueryStrParam(dc, "", "", "", filter, ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}

	httpClient := http_client.HttpClient{
		Method:      health.NODE_HEALTH_STATUS[0],
		Url:         url,
		ContentType: health.NODE_HEALTH_STATUS[2],
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
	var statusArr []health.NodeStatus
	err = json.Unmarshal(jsonBytes, &statusArr)
	if err != nil {
		return nil, err
	}
	return statusArr, nil
}

func (client *ConsulClient) ServiceHealthStatus(service, ns, dc, filter, near, nodeMeta string) ([]health.ServiceStatus, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(health.SERVICE_HEALTH_STATUS[1]), service)
	paras := packageQueryStrParam(dc, "", near, nodeMeta, filter, ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}

	httpClient := http_client.HttpClient{
		Method:      health.SERVICE_HEALTH_STATUS[0],
		Url:         url,
		ContentType: health.SERVICE_HEALTH_STATUS[2],
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
	var statusArr []health.ServiceStatus
	err = json.Unmarshal(jsonBytes, &statusArr)
	if err != nil {
		return nil, err
	}
	return statusArr, nil
}

func (client *ConsulClient) ServiceInstances(service, ns, dc, filter, near, tag, nodeMeta string, passing bool) ([]health.ServiceInstance, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(health.SERVICE_INSTANCES[1]), service)
	paras := packageQueryStrParam(dc, tag, near, nodeMeta, filter, ns, "", "", "")
	if len(paras) != 0 {
		paras = fmt.Sprintf("%s&%s", paras, packageQueryBoolParam(false, false, false, false, passing, false))
		url = fmt.Sprintf("%s?%s", url, paras)
	}

	httpClient := http_client.HttpClient{
		Method:      health.SERVICE_INSTANCES[0],
		Url:         url,
		ContentType: health.SERVICE_INSTANCES[2],
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
	var instances []health.ServiceInstance
	err = json.Unmarshal(jsonBytes, &instances)
	if err != nil {
		return nil, err
	}
	return instances, nil
}
