package consul_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jianghaibo12138/go-consul/consul_client/check"
	"github.com/jianghaibo12138/go-consul/http_client"
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
	reqBytes, err := json.Marshal([]byte{})
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
