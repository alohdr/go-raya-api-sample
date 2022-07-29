package server

import (
	"fmt"
	"go-rest-api/app"
	"go-rest-api/config"
	"net/http"
)

func NewServer() {
	cfg := config.Get().Port
	host := fmt.Sprintf(":%v", cfg)

	server := new(http.Server)
	server.Addr = host

	//var redisHost = "localhost:6379"
	//var redisPassword = ""
	//
	//rdb := redis.NewRedisClient(redisHost, redisPassword)
	//fmt.Println("redis client initialized ", rdb)

	appService := app.NewService()
	app.NewRouter(appService)

	fmt.Printf("Server Running on locahost%v", host)

	server.ListenAndServe()

}
