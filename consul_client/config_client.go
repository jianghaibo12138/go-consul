package consul_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jianghaibo12138/go-consul/consul_client/config"
	"github.com/jianghaibo12138/go-consul/http_client"
	"io/ioutil"
)

func (client *ConsulClient) ConfigUpsert(dc, ns, kind, name, protocol string, cas int) (bool, error) {
	url := client.packageRequestTpl(config.CONFIG_UPSERT[1])
	paras := packageQueryStrParam(dc, "", "", "", "", ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s&cas=%d", url, paras, cas)
	}
	jsonBytes, err := json.Marshal(config.PayloadConfig{
		Kind:     kind,
		Name:     name,
		Protocol: protocol,
	})
	httpClient := http_client.HttpClient{
		Method:      config.CONFIG_UPSERT[0],
		Url:         url,
		ContentType: config.CONFIG_UPSERT[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request(jsonBytes)
	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		return false, errors.New(response.Status)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	var success bool
	err = json.Unmarshal(bytes, &success)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (client *ConsulClient) ConfigGetter(dc, ns, kind, name string) (*config.ResponseConfig, error) {
	url := fmt.Sprintf("%s/%s/%s", client.packageRequestTpl(config.CONFIG_GETTER[1]), kind, name)
	paras := packageQueryStrParam(dc, "", "", "", "", ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      config.CONFIG_GETTER[0],
		Url:         url,
		ContentType: config.CONFIG_GETTER[2],
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
	var res config.ResponseConfig
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (client *ConsulClient) ConfigListGetter(dc, ns, kind string) (*[]config.ResponseConfig, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(config.CONFIG_LIST_GETTER[1]), kind)
	paras := packageQueryStrParam(dc, "", "", "", "", ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      config.CONFIG_LIST_GETTER[0],
		Url:         url,
		ContentType: config.CONFIG_LIST_GETTER[2],
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
	var res []config.ResponseConfig
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (client *ConsulClient) ConfigDelete(dc, ns, kind, name string) (*config.ResponseConfig, error) {
	url := fmt.Sprintf("%s/%s/%s", client.packageRequestTpl(config.CONFIG_DELETE[1]), kind, name)
	paras := packageQueryStrParam(dc, "", "", "", "", ns, "", "", "")
	if len(paras) != 0 {
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	httpClient := http_client.HttpClient{
		Method:      config.CONFIG_DELETE[0],
		Url:         url,
		ContentType: config.CONFIG_DELETE[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return nil, nil
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, nil
	}
	return nil, nil
}
