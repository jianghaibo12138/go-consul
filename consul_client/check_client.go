package consul_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jianghaibo12138/go-consul/consul_client/check"
	"github.com/jianghaibo12138/go-consul/http_client"
	"io/ioutil"
)

// https://www.consul.io/api-docs/agent/check

//
// CheckRegister
// @Description: 注册健康检查服务
// @receiver client
// @param payload
// @return bool
// @return error
//
func (client *ConsulClient) CheckRegister(payload check.RegisterPayload) (bool, error) {
	httpClient := http_client.HttpClient{
		Method:      check.CHECK_REGISTER[0],
		Url:         client.packageRequestTpl(check.CHECK_REGISTER[1]),
		ContentType: check.CHECK_REGISTER[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	reqBytes, err := json.Marshal(payload)
	if err != nil {
		return false, err
	}
	response, err := httpClient.Request(reqBytes)
	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		return false, errors.New(response.Status)
	}
	return true, nil
}

//
// CheckDeRegister
// @Description: 注销活跃检查服务
// @receiver client
// @param checkId
// @return bool
// @return error
//
func (client *ConsulClient) CheckDeRegister(checkId string) (bool, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(check.CHECK_DEREGISTER[1]), checkId)
	httpClient := http_client.HttpClient{
		Method:      check.CHECK_DEREGISTER[0],
		Url:         url,
		ContentType: check.CHECK_DEREGISTER[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		return false, errors.New(response.Status)
	}
	return true, nil
}

//
// CheckList
// @Description: 获取所有的检查服务
// @receiver client
// @param filter
// @param ns
// @return map[string]interface{}
// @return error
//
func (client *ConsulClient) CheckList(filter, ns string) (map[string]interface{}, error) {
	url := client.packageRequestTpl(check.CHECK_LIST[1])
	params := packageQueryStrParam("", "", "", "", filter, ns, "", "", "", "", "", "", "")
	if len(params) != 0 {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	httpClient := http_client.HttpClient{
		Method:      check.CHECK_LIST[0],
		Url:         url,
		ContentType: check.CHECK_LIST[2],
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
	var m map[string]interface{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

