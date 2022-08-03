package server

import (
	"fmt"
	"go-rest-api/app"
	"go-rest-api/auth"
	"go-rest-api/config"
	"net/http"
)

func NewServer() {
	cfg := config.Get().Port
	host := fmt.Sprintf(":%v", cfg)

	server := new(http.Server)
	server.Addr = host

	appService := app.NewService()
	authService := auth.NewService()
	app.NewRouter(appService)
	auth.NewRouter(authService)

	fmt.Printf("Server Running on locahost%v", host)

	server.ListenAndServe()

}
