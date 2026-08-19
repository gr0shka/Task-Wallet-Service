package main

import (
	server "Task-Wallet-Service/internal/delivery/http"
	"Task-Wallet-Service/internal/domain"
	repository "Task-Wallet-Service/internal/repository/memory"
	service "Task-Wallet-Service/internal/service"
	"context"
)

func main() {
	rep := repository.NewInMemoryUserRepository()
	rep.AddUser(context.Background(), &domain.User{1, 1000})

	sv := service.NewUserService(rep)
	sr := server.NewUserHandler(sv)

	sr.StartListen()
}
