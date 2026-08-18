package http

import (
	"Task-Wallet-Service/internal/domain"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type UserHandler struct {
	sv domain.UserService
}

func NewUserHandler(sv domain.UserService) *UserHandler {
}

func (u *UserHandler) DepositHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user, err := u.sv.Deposit(context.Background(), id, amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", user)))
	return
}
