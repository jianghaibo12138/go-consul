package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/go-consul/consul_client"
	"github.com/jianghaibo12138/go-consul/consul_client/service"
	"os"
)

const (
	Host  = "127.0.0.1"
	Port  = 8500
	Token = "e95b599e-166e-7d80-08ad-aee76e7ddf19"
	Ssl   = false
)

func consulRegister() {
	fmt.Println(os.Getwd())
	client := consul_client.ConsulClient{
		Host:  Host,
		Port:  Port,
		Token: Token,
		Ssl:   Ssl,
	}
	bytes := consul_client.ReadJsonConf("./examples/gin-consul/json_conf/service_register.json")
	var srv service.RegisterService
	err := json.Unmarshal(bytes, &srv)
	if err != nil {
		panic(err)
	}
	infoS, err := client.RegisterService(srv)
	if err != nil {
		panic(err)
	}
	fmt.Println(infoS)
}

func main() {
	consulRegister()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(err)
	}
}
