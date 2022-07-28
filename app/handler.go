package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) Handler {
	return Handler{
		service: s,
	}
}

//USE
func (h Handler) handleGetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "GET" {

		response, err := h.service.GetAllUser(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}
func (h Handler) handleInsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "POST" {

		u, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// var user User
		// err = json.Unmarshal(u, &user)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }

		response, errResp := h.service.InsertUser(ctx, u)
		if errResp != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, errResult := json.Marshal(response)
		if errResult != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}

//ACCOUNT
func (h Handler) handleGetAllAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "GET" {
		response, err := h.service.GetAllAccount(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}

//BANK
func (h Handler) handleGetAllBank(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "GET" {

		filter := r.FormValue("bank_name")
		fmt.Println(filter)
		response, err := h.service.GetAllBank(ctx, filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}
func (h Handler) handleInsertBank(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "POST" {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var bank BankCollect
		err = json.Unmarshal(b, &bank)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		response, err := h.service.InsertBank(ctx, bank)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}

//TRANSACTION
func (h Handler) handleGetLastTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "GET" {

		filterUser := r.FormValue("user")
		filterBank := r.FormValue("bank")

		response, err := h.service.GetLastTransaction(ctx, filterUser, filterBank)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}
func (h Handler) handleGetFavoriteTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "GET" {

		//filter := r.FormValue("bank_name")

		response, err := h.service.GetFavoriteTransaction(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}
func (h Handler) handleCheckAccountNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "GET" {

		filter := r.FormValue("account_bank_number")
		fmt.Println(filter)
		response, err := h.service.CheckAccountNumber(ctx, filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}
func (h Handler) handleGetSaldo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "GET" {

		filterUser := r.FormValue("user")
		filterPin := r.FormValue("pin")

		fmt.Println(filterUser)
		fmt.Println(filterPin)

		response, err := h.service.GetSaldo(ctx, filterUser, filterPin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}
func (h Handler) handleGetTransactionDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "GET" {

		id := r.FormValue("transaction_id")

		response, err := h.service.GetTransactionDetail(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}
func (h Handler) handleUpdateStatusTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "PUT" {

		status, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, errResp := h.service.UpdateStatusTransaction(ctx, status)
		if errResp != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, errResult := json.Marshal(response)
		if errResult != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}
func (h Handler) handleInsertTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "POST" {

		tf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// var user User
		// err = json.Unmarshal(u, &user)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }

		response, errResp := h.service.InsertTransaction(ctx, tf)
		if errResp != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, errResult := json.Marshal(response)
		if errResult != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}
