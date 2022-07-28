package app

import (
	"net/http"
)

func NewRouter(service Service) {
	handler := NewHandler(service)

	//USER
	http.HandleFunc("/users", handler.handleGetAllUser)
	http.HandleFunc("/user", handler.handleInsertUser)
	http.HandleFunc("/saldo", handler.handleGetSaldo)

	//ACCOUNT
	http.HandleFunc("/accounts", handler.handleGetAllAccount)
	http.HandleFunc("/accounts/check", handler.handleCheckAccountNumber)

	//BANK
	http.HandleFunc("/bank", handler.handleGetAllBank)
	http.HandleFunc("/bank/create", handler.handleInsertBank)

	//TRANSACTION
	http.HandleFunc("/transactions", handler.handleGetLastTransaction)
	http.HandleFunc("/transactions/favorite", handler.handleGetFavoriteTransaction)
	http.HandleFunc("/transactions/detail", handler.handleGetTransactionDetail)
	http.HandleFunc("/transactions/update", handler.handleUpdateStatusTransaction)
	http.HandleFunc("/transactions/create", handler.handleInsertTransaction)
}
