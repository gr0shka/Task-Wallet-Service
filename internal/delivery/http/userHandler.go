package http

import (
	"Task-Wallet-Service/internal/domain"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type UserHandler struct {
	sv domain.WalletUsecase
}

type DepositRequest struct {
	UserId int64 `json:"user_id"`
	Amount int64 `json:"amount"`
}

func NewUserHandler(sv domain.WalletUsecase) *UserHandler {
	return &UserHandler{sv: sv}
}

func (u *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/wallet/deposit", u.DepositHandler)
}

func (u *UserHandler) StartListen(mux *http.ServeMux) error {
	log.Println("Starting HTTP server on port 8080")
	return http.ListenAndServe(":8080", mux)
}

func (u *UserHandler) DepositHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req DepositRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updatedUser, err := u.sv.Deposit(r.Context(), req.UserId, req.Amount)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(updatedUser)
}
