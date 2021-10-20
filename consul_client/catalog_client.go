package consul_client

import (
	"amazing2j.com/go-consul/consul_client/catalog"
	"amazing2j.com/go-consul/http_client"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func packageQueryStrParam(dc, tag, near, nodeMeta, filter, ns, acquire, release string) string {
	var paraStrArr []string
	if dc != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("dc=%s", dc))
	}
	if tag != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("tag=%s", tag))
	}
	if near != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("near=%s", near))
	}
	if nodeMeta != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("node-meta=%s", nodeMeta))
	}
	if filter != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("filter=%s", filter))
	}
	if ns != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("ns=%s", ns))
	}
	if acquire != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("acquire=%s", acquire))
	}
	if release != "" {
		paraStrArr = append(paraStrArr, fmt.Sprintf("release=%s", release))
	}
	var paraStr string
	if len(paraStrArr) != 0 {
		paraStr = strings.Join(paraStrArr, "&")
	}
	return paraStr
}

func packageQueryBoolParam(recurse, raw, keys, separator bool) string {
	var paraStrArr []string
	if recurse {
		paraStrArr = append(paraStrArr, fmt.Sprintf("recurse=%v", recurse))
	}
	if raw {
		paraStrArr = append(paraStrArr, fmt.Sprintf("raw=%v", raw))
	}
	if keys {
		paraStrArr = append(paraStrArr, fmt.Sprintf("keys=%v", keys))
	}
	if separator {
		paraStrArr = append(paraStrArr, fmt.Sprintf("separator=%v", separator))
	}
	var paraStr string
	if len(paraStrArr) != 0 {
		paraStr = strings.Join(paraStrArr, "&")
	}
	return paraStr
}

func packageQueryIntParam(flags, cas int) string {
	var paraStrArr []string
	if flags != 0 {
		paraStrArr = append(paraStrArr, fmt.Sprintf("flags=%d", flags))
	}
	if cas != 0 {
		paraStrArr = append(paraStrArr, fmt.Sprintf("cas=%d", cas))
	}
	var paraStr string
	if len(paraStrArr) != 0 {
		paraStr = strings.Join(paraStrArr, "&")
	}
	return paraStr
}

func (client *ConsulClient) CatalogRegister(namespace string, cal *catalog.RegisterCatalog) (*[]byte, error) {
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_REGISTER[0],
		Url:         fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, catalog.CATALOG_REGISTER[1]),
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
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &bytes, nil
}

func (client *ConsulClient) CatalogDeRegister(namespace string, cal *catalog.DeRegisterCatalog) (*[]byte, error) {
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_DEREGISTER[0],
		Url:         fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, catalog.CATALOG_DEREGISTER[1]),
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
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &bytes, nil
}

func (client *ConsulClient) CatalogDatacenters() ([]string, error) {
	httpClient := http_client.HttpClient{
		Method:      catalog.CATALOG_DATACENTERS[0],
		Url:         fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, catalog.CATALOG_DATACENTERS[1]),
		ContentType: catalog.CATALOG_DATACENTERS[2],
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
	var dcs []string
	err = json.Unmarshal(bytes, &dcs)
	if err != nil {
		return nil, err
	}
	return dcs, nil
}

func (client *ConsulClient) CatalogNodes(dc, near, filter string) ([]catalog.Node, error) {
	url := fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, catalog.CATALOG_NODES[1])
	paras := packageQueryStrParam(dc, "", near, "", filter, "", "", "")
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
	url := fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, catalog.CATALOG_SERVICES[1])
	paras := packageQueryStrParam(dc, "", "", nodeMeta, "", namespace, "", "")
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
	url := fmt.Sprintf("%s/%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, catalog.CATALOG_SERVICE_NODES[1]), service)
	paras := packageQueryStrParam(dc, tag, near, "", filter, ns, "", "")
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
	url := fmt.Sprintf("%s/%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, catalog.CATALOG_CONNECT_NODES[1]), service)
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
	url := fmt.Sprintf("%s/%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, catalog.CATALOG_NODE_NODES[1]), node)
	paras := packageQueryStrParam(dc, "", "", "", filter, ns, "", "")
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
	url := fmt.Sprintf("%s/%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, catalog.CATALOG_NODE_SERVICES[1]), node)
	paras := packageQueryStrParam(dc, "", "", "", filter, ns, "", "")
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
	url := fmt.Sprintf("%s/%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, catalog.CATALOG_GATEWAY_SERVICES[1]), gateway)
	paras := packageQueryStrParam(dc, "", "", "", "", ns, "", "")
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
