package consul_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jianghaibo12138/go-consul/consul_client/kv"
	"jianghaibo12138/go-consul/http_client"
)

func (client *ConsulClient) UpsertKey(ns, key, dc, acquire, release string, flags, cas int, value []byte) (bool, error) {
	url := fmt.Sprintf("%s/%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, kv.KV_UPSERT_KEY[1]), key)
	paras := packageQueryStrParam(dc, "", "", "", "", ns, acquire, release)
	if len(paras) != 0 {
		paras = fmt.Sprintf("%s&%s", paras, packageQueryIntParam(flags, cas))
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	parasInt := packageQueryIntParam(flags, cas)
	if len(parasInt) != 0 {
		url = fmt.Sprintf("%s?%s", url, parasInt)
	}

	httpClient := http_client.HttpClient{
		Method:      kv.KV_UPSERT_KEY[0],
		Url:         url,
		ContentType: kv.KV_UPSERT_KEY[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request(value)
	if err != nil {
		return false, err
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	var success bool
	err = json.Unmarshal(jsonBytes, &success)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (client *ConsulClient) ReadKey(ns, key, dc string, recurse, raw, keys, separator bool) (interface{}, error) {
	url := fmt.Sprintf("%s/%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, kv.KV_READ_KEY[1]), key)
	paras := packageQueryStrParam(dc, "", "", "", "", ns, "", "")
	if len(paras) != 0 {
		paras = fmt.Sprintf("%s&%s", paras, packageQueryBoolParam(recurse, raw, keys, separator))
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	parasBool := packageQueryBoolParam(recurse, raw, keys, separator)
	if len(parasBool) != 0 {
		url = fmt.Sprintf("%s?%s", url, parasBool)
	}

	httpClient := http_client.HttpClient{
		Method:      kv.KV_READ_KEY[0],
		Url:         url,
		ContentType: kv.KV_READ_KEY[2],
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
	if keys {
		var kvInfo []string
		err = json.Unmarshal(jsonBytes, &kvInfo)
		if err != nil {
			return nil, err
		}
		return &kvInfo, nil
	}
	if raw {
		return string(jsonBytes), nil
	}
	var kvInfo []kv.KeyValue
	err = json.Unmarshal(jsonBytes, &kvInfo)
	if err != nil {
		return nil, err
	}
	return &kvInfo, nil
}

func (client *ConsulClient) DeleteKey(ns, key, dc string, cas int, recurse bool) (bool, error) {
	url := fmt.Sprintf("%s/%s", fmt.Sprintf(client.packageRequestTpl(), client.Host, client.Port, kv.KV_DELETE_KEY[1]), key)
	paras := packageQueryStrParam(dc, "", "", "", "", ns, "", "")
	if len(paras) != 0 {
		paras = fmt.Sprintf("%s&%s", paras, packageQueryIntParam(0, cas))
		url = fmt.Sprintf("%s?%s", url, paras)
	}
	parasInt := packageQueryIntParam(0, cas)
	if len(parasInt) != 0 {
		url = fmt.Sprintf("%s?%s", url, parasInt)
	}
	parasBool := packageQueryBoolParam(recurse, false, false, false)
	if len(parasBool) != 0 {
		url = fmt.Sprintf("%s?%s", url, parasBool)
	}

	httpClient := http_client.HttpClient{
		Method:      kv.KV_DELETE_KEY[0],
		Url:         url,
		ContentType: kv.KV_DELETE_KEY[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return false, err
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	var success bool
	err = json.Unmarshal(jsonBytes, &success)
	if err != nil {
		return false, err
	}
	return success, nil
}
