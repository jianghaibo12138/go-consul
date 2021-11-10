package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jianghaibo12138/go-consul/consul_client"
	"github.com/jianghaibo12138/go-consul/consul_client/check"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	Host  = "127.0.0.1"
	Port  = 8500
	Token = "e95b599e-166e-7d80-08ad-aee76e7ddf19"
	Ssl   = false
)

var client = consul_client.ConsulClient{
	Host:  Host,
	Port:  Port,
	Token: Token,
	Ssl:   Ssl,
}

func consulRegister() {
	bytes := consul_client.ReadJsonConf("./examples/gin-consul/json_conf/check_register.json")
	var payload check.RegisterPayload
	err := json.Unmarshal(bytes, &payload)
	if err != nil {
		panic(err)
	}
	infoS, err := client.CheckRegister(payload)
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

	srv := &http.Server{
		Addr:    "0.0.0:8080",
		Handler: r,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(fmt.Sprintf("listen: %s", err))
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server")

	_, err := client.CheckDeRegister("gin-consul")
	if err != nil {
		panic(err)
	}

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(fmt.Sprintf("Server forced to shutdown: %s", err.Error()))
	}
	log.Println("Server exiting")
}
