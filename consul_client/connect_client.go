package consul_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jianghaibo12138/go-consul/consul_client/connect"
	"github.com/jianghaibo12138/go-consul/http_client"
	"io/ioutil"
)

// https://www.consul.io/api-docs/agent/connect

//
// GetConnectCARoots
// @Description: 获取CA Roots
// @receiver client
// @return *connect.CADetail
// @return error
//
func (client *ConsulClient) GetConnectCARoots() (*connect.CADetail, error) {
	httpClient := http_client.HttpClient{
		Method:      connect.GET_CONNECT_CA_ROOT[0],
		Url:         client.packageRequestTpl(connect.GET_CONNECT_CA_ROOT[1]),
		ContentType: connect.GET_CONNECT_CA_ROOT[2],
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
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	var caDetail connect.CADetail
	err = json.Unmarshal(bytes, &caDetail)
	if err != nil {
		return nil, err
	}
	return &caDetail, nil
}

//
// GetConnectCAConfiguration
// @Description: 获取CA配置数据
// @receiver client
// @return *connect.CAConfiguration
// @return error
//
func (client *ConsulClient) GetConnectCAConfiguration() (*connect.CAConfiguration, error) {
	httpClient := http_client.HttpClient{
		Method:      connect.GET_CONNECT_CA_CONFIGURATION[0],
		Url:         client.packageRequestTpl(connect.GET_CONNECT_CA_CONFIGURATION[1]),
		ContentType: connect.GET_CONNECT_CA_CONFIGURATION[2],
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
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	var caConfigs connect.CAConfiguration
	err = json.Unmarshal(bytes, &caConfigs)
	if err != nil {
		return nil, err
	}
	return &caConfigs, nil
}

//
// UpdateConnectCAConfiguration
// @Description: 更新CA配置数据
// @receiver client
// @param config
// @return error
//
func (client *ConsulClient) UpdateConnectCAConfiguration(config *connect.CAConfiguration) error {
	httpClient := http_client.HttpClient{
		Method:      connect.UPDATE_CONNECT_CA_CONFIGURATION[0],
		Url:         client.packageRequestTpl(connect.UPDATE_CONNECT_CA_CONFIGURATION[1]),
		ContentType: connect.UPDATE_CONNECT_CA_CONFIGURATION[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	reqBytes, err := json.Marshal(config)
	if err != nil {
		return err
	}
	response, err := httpClient.Request(reqBytes)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	return nil
}

//
// GetIntentionByName
// @Description: 通过name获取"意图"数据
// @receiver client
// @param source
// @param destination
// @return *connect.IntentionPayload
// @return error
//
func (client *ConsulClient) GetIntentionByName(source, destination string) (*connect.IntentionPayload, error) {
	url := client.packageRequestTpl(connect.GET_INTENTION_BY_NAME[1])
	params := packageQueryStrParam("", "", "", "", "", "", "", "", "", source, destination, "", "")
	if len(params) != 0 {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	httpClient := http_client.HttpClient{
		Method:      connect.GET_INTENTION_BY_NAME[0],
		Url:         url,
		ContentType: connect.GET_INTENTION_BY_NAME[2],
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
	var resp connect.IntentionPayload
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	return &resp, nil
}

//
// GetIntentionById
// @Description: 通过Id获取"意图"数据
// @receiver client
// @param id
// @return *connect.IntentionPayload
// @return error
//
func (client *ConsulClient) GetIntentionById(id string) (*connect.IntentionPayload, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(connect.GET_INTENTION_BY_ID[1]), id)
	httpClient := http_client.HttpClient{
		Method:      connect.GET_INTENTION_BY_ID[0],
		Url:         url,
		ContentType: connect.GET_INTENTION_BY_ID[2],
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
	var resp connect.IntentionPayload
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	return &resp, nil
}

//
// GetIntentionList
// @Description: 获取所有"意图"数据
// @receiver client
// @return *[]connect.IntentionPayload
// @return error
//
func (client *ConsulClient) GetIntentionList(filter string) (*[]connect.IntentionPayload, error) {
	url := client.packageRequestTpl(connect.GET_INTENTION_BY_ID[1])
	params := packageQueryStrParam("", "", "", "", filter, "", "", "", "", "", "", "", "")
	if len(params) != 0 {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	httpClient := http_client.HttpClient{
		Method:      connect.GET_INTENTION_BY_ID[0],
		Url:         url,
		ContentType: connect.GET_INTENTION_BY_ID[2],
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
	var resp []connect.IntentionPayload
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	return &resp, nil
}

//
// UpsertIntentionByName
// @Description: 通过名称更新或新建"意图"
// @receiver client
// @param sourceService
// @param destinationService
// @param payload
// @return bool
// @return error
//
func (client *ConsulClient) UpsertIntentionByName(source, destination string, payload *connect.IntentionPayload) (bool, error) {
	url := client.packageRequestTpl(connect.UPSERT_INTENTIONS_BY_NAME[1])
	params := packageQueryStrParam("", "", "", "", "", "", "", "", "", source, destination, "", "")
	if len(params) != 0 {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	httpClient := http_client.HttpClient{
		Method:      connect.UPSERT_INTENTIONS_BY_NAME[0],
		Url:         url,
		ContentType: connect.UPSERT_INTENTIONS_BY_NAME[2],
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
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	var success bool
	err = json.Unmarshal(bytes, &success)
	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		return false, errors.New(response.Status)
	}
	return success, nil
}

//
// CreateIntentionWithId
// @Description: 新建"意图"并返回唯一ID
// @receiver client
// @param payload
// @return *connect.IntentionResponse
// @return error
//
func (client *ConsulClient) CreateIntentionWithId(payload *connect.IntentionPayload) (*connect.IntentionResponse, error) {
	url := client.packageRequestTpl(connect.CREATE_INTENTIONS_WITH_ID[1])
	httpClient := http_client.HttpClient{
		Method:      connect.CREATE_INTENTIONS_WITH_ID[0],
		Url:         url,
		ContentType: connect.CREATE_INTENTIONS_WITH_ID[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	reqBytes, err := json.Marshal(payload)
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
	var resp connect.IntentionResponse
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	return &resp, nil
}

//
// UpdateIntentionById
// @Description: 通过Id更新"意图"数据
// @receiver client
// @param id
// @param payload
// @return *connect.IntentionResponse
// @return error
//
func (client *ConsulClient) UpdateIntentionById(id string, payload *connect.IntentionPayload) (*connect.IntentionResponse, error) {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(connect.UPDATE_INTENTIONS_BY_ID[1]), id)
	httpClient := http_client.HttpClient{
		Method:      connect.UPDATE_INTENTIONS_BY_ID[0],
		Url:         url,
		ContentType: connect.UPDATE_INTENTIONS_BY_ID[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	reqBytes, err := json.Marshal(payload)
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
	var resp connect.IntentionResponse
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	return &resp, nil
}

//
// DeleteIntentionByName
// @Description: 通过名称删除"意图"数据
// @receiver client
// @param source
// @param destination
// @return *connect.IntentionResponse
// @return error
//
func (client *ConsulClient) DeleteIntentionByName(source, destination string) error {
	url := client.packageRequestTpl(connect.DELETE_INTENTION_BY_NAME[1])
	params := packageQueryStrParam("", "", "", "", "", "", "", "", "", source, destination, "", "")
	if len(params) != 0 {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	httpClient := http_client.HttpClient{
		Method:      connect.DELETE_INTENTION_BY_NAME[0],
		Url:         url,
		ContentType: connect.DELETE_INTENTION_BY_NAME[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	return nil
}

//
// DeleteIntentionById
// @Description: 通过ID删除"意图"数据
// @receiver client
// @param id
// @return error
//
func (client *ConsulClient) DeleteIntentionById(id string) error {
	url := fmt.Sprintf("%s/%s", client.packageRequestTpl(connect.DELETE_INTENTION_BY_ID[1]), id)
	httpClient := http_client.HttpClient{
		Method:      connect.DELETE_INTENTION_BY_ID[0],
		Url:         url,
		ContentType: connect.DELETE_INTENTION_BY_ID[2],
		Headers:     map[string]string{CONSUL_TOKEN_KEY: client.Token},
	}
	response, err := httpClient.Request([]byte{})
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return errors.New(response.Status)
	}
	return nil
}

//
// CheckIntentionResult
// @Description: 检查"意图"执行结果数据
// @receiver client
// @param source
// @param destination
// @return *connect.IntentionResponse
// @return error
//
func (client *ConsulClient) CheckIntentionResult(source, destination string) (*connect.IntentionResponse, error) {
	url := client.packageRequestTpl(connect.CHECK_INTENTION_RESULT[1])
	params := packageQueryStrParam("", "", "", "", "", "", "", "", "", source, destination, "", "")
	if len(params) != 0 {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	httpClient := http_client.HttpClient{
		Method:      connect.CHECK_INTENTION_RESULT[0],
		Url:         url,
		ContentType: connect.CHECK_INTENTION_RESULT[2],
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
	var resp connect.IntentionResponse
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (client *ConsulClient) GetMatchingIntentionList(by, name string) (*map[string][]connect.IntentionPayload, error) {
	url := client.packageRequestTpl(connect.CHECK_INTENTION_RESULT[1])
	params := packageQueryStrParam("", "", "", "", "", "", "", "", "", "", "", by, name)
	if len(params) != 0 {
		url = fmt.Sprintf("%s?%s", url, params)
	}
	httpClient := http_client.HttpClient{
		Method:      connect.CHECK_INTENTION_RESULT[0],
		Url:         url,
		ContentType: connect.CHECK_INTENTION_RESULT[2],
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
	var resp map[string][]connect.IntentionPayload
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
