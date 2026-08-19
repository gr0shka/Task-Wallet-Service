package main

import (
	server "Task-Wallet-Service/internal/delivery/http"
	"Task-Wallet-Service/internal/domain"
	repository "Task-Wallet-Service/internal/repository/memory"
	"Task-Wallet-Service/internal/usecase"
	"context"
	"net/http"
)

func main() {
	rep := repository.NewInMemoryUserRepository()
	rep.AddUser(context.Background(), &domain.User{1, 1000})

	walletUsecase := usecase.NewWalletUsecase(rep)
	sr := server.NewUserHandler(walletUsecase)
	mux := http.NewServeMux()
	sr.RegisterRoutes(mux)

	_ = sr.StartListen(mux)
}
