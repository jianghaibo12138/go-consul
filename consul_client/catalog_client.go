package consul_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jianghaibo12138/go-consul/consul_client/catalog"
	"github.com/jianghaibo12138/go-consul/http_client"
	"io/ioutil"
)

func (client *ConsulClient) CatalogRegister(namespace string, cal *catalog.RegisterCatalog) (*[]byte, error) {
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_REGISTER[0],
		Url:         client.packageRequestTpl(catalog.CATALOG_REGISTER[1]),
		ContentType: catalog.CATALOG_REGISTER[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token, CONSUL_NAMESPACE_KEY: namespace},
	}
	reqBytes, err := json.Marshal(cal)
	if err != nil {
		return nil, err
	}
	response, err := httpClient.Request(reqBytes)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &bytes, nil
}

func (client *ConsulClient) CatalogDeRegister(namespace string, cal *catalog.DeRegisterCatalog) (*[]byte, error) {
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_DEREGISTER[0],
		Url:         client.packageRequestTpl(catalog.CATALOG_DEREGISTER[1]),
		ContentType: catalog.CATALOG_DEREGISTER[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token, CONSUL_NAMESPACE_KEY: namespace},
	}
	reqBytes, err := json.Marshal(cal)
	if err != nil {
		return nil, err
	}
	response, err := httpClient.Request(reqBytes)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &bytes, nil
}

func (client *ConsulClient) CatalogDatacenters() ([]string, error) {
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_DATACENTERS[0],
		Url:         client.packageRequestTpl(catalog.CATALOG_DATACENTERS[1]),
		ContentType: catalog.CATALOG_DATACENTERS[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var dcs []string
	err = json.Unmarshal(bytes, &dcs)
	if err != nil {
		return nil, err
	}
	return dcs, nil
}

func (client *ConsulClient) CatalogNodes(dc, near, filter string) ([]catalog.Node, error) {
	url := client.packageRequestTpl(catalog.CATALOG_NODES[1])
	paras := packageQueryStrParam(dc, "", near, "", filter, "", "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_NODES[0],
		Url:         url,
		ContentType: catalog.CATALOG_NODES[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var nodes []catalog.Node
	err = json.Unmarshal(bytes, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (client *ConsulClient) CatalogServices(dc, nodeMeta, namespace string) (map[string][]string, error) {
	url := client.packageRequestTpl(catalog.CATALOG_SERVICES[1])
	paras := packageQueryStrParam(dc, "", "", nodeMeta, "", namespace, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_SERVICES[0],
		Url:         url,
		ContentType: catalog.CATALOG_SERVICES[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var services = make(map[string][]string)
	err = json.Unmarshal(bytes, &services)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (client *ConsulClient) CatalogServiceNodes(service, dc, tag, near, filter, ns string) ([]catalog.Node, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(catalog.CATALOG_SERVICE_NODES[1]), service)
	paras := packageQueryStrParam(dc, tag, near, "", filter, ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_SERVICE_NODES[0],
		Url:         url,
		ContentType: catalog.CATALOG_SERVICE_NODES[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var nodes []catalog.Node
	err = json.Unmarshal(bytes, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (client *ConsulClient) CatalogConnectNodes(service string) ([]catalog.Node, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(catalog.CATALOG_CONNECT_NODES[1]), service)
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_CONNECT_NODES[0],
		Url:         url,
		ContentType: catalog.CATALOG_CONNECT_NODES[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var nodes []catalog.Node
	err = json.Unmarshal(bytes, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (client *ConsulClient) CatalogNodeNodes(node, dc, filter, ns string) ([]catalog.Node, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(catalog.CATALOG_NODE_NODES[1]), node)
	paras := packageQueryStrParam(dc, "", "", "", filter, ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_NODE_NODES[0],
		Url:         url,
		ContentType: catalog.CATALOG_NODE_NODES[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var nodes []catalog.Node
	err = json.Unmarshal(bytes, &nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func (client *ConsulClient) CatalogNodeServices(node, dc, filter, ns string) (*catalog.NodeServices, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(catalog.CATALOG_NODE_SERVICES[1]), node)
	paras := packageQueryStrParam(dc, "", "", "", filter, ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_NODE_SERVICES[0],
		Url:         url,
		ContentType: catalog.CATALOG_NODE_SERVICES[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var services catalog.NodeServices
	err = json.Unmarshal(bytes, &services)
	if err != nil {
		return nil, err
	}
	return &services, nil
}

func (client *ConsulClient) CatalogGatewayServices(gateway, dc, ns string) (*catalog.Gateway, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(catalog.CATALOG_GATEWAY_SERVICES[1]), gateway)
	paras := packageQueryStrParam(dc, "", "", "", "", ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_GATEWAY_SERVICES[0],
		Url:         url,
		ContentType: catalog.CATALOG_GATEWAY_SERVICES[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var gateways catalog.Gateway
	err = json.Unmarshal(bytes, &gateways)
	if err != nil {
		return nil, err
	}
	return &gateways, nil
}
