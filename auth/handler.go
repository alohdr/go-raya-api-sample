package auth

import (
	"context"
	"encoding/json"
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

func (h Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx context.Context = context.Background()

	if r.Method == "POST" {

		u, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var user LoginParams
		err = json.Unmarshal(u, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		res, err := h.service.Login(ctx, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		result, errResult := json.Marshal(res)
		if errResult != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	}

	http.Error(w, "", http.StatusBadRequest)
}
