package http_client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jianghaibo12138/go-consul/consul_client/agent"
	"io"
	"net/http"
)

type HttpClient struct {
	Method      string            `json:"method"`
	Url         string            `json:"url"`
	ContentType string            `json:"content_type"`
	Headers     map[string]string `json:"headers"`
}

func (client *HttpClient) Request(data []byte) (*http.Response, error) {
	reader := bytes.NewReader(data)
	request, err := http.NewRequest(client.Method, client.Url, reader)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = request.Body.Close()
	}() // 程序在使用完回复后必须关闭回复的主体
	if len(client.ContentType) != 0 {
		request.Header.Set("Content-Type", client.ContentType)
	} else {
		request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	}
	if len(client.Headers) != 0 {
		for key, value := range client.Headers {
			request.Header.Set(key, value)
		}
	}
	c := http.Client{}
	resp, err := c.Do(request) // Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (client *HttpClient) RequestStream(data []byte) (*http.Response, error) {
	// TODO 存在流式响应读取问题
	request, err := http.NewRequest(client.Method, client.Url, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = request.Body.Close()
	}() // 程序在使用完回复后必须关闭回复的主体
	if len(client.ContentType) != 0 {
		request.Header.Set("Content-Type", client.ContentType)
	} else {
		request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	}
	if len(client.Headers) != 0 {
		for key, value := range client.Headers {
			request.Header.Set(key, value)
		}
	}
	resp, err := http.DefaultClient.Do(request) // Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		var logMsg agent.StreamLog
		err = json.Unmarshal(line, &logMsg)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		fmt.Println(fmt.Sprintf("%+v", logMsg))
	}
	return resp, nil
}
