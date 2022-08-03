package auth

import "net/http"

func NewRouter(service Service) {
	handler := NewHandler(service)

	//LOGIN
	http.HandleFunc("/login", handler.handleLogin)
}
