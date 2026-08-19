package main

import (
	server "Task-Wallet-Service/internal/delivery/http"
	repository "Task-Wallet-Service/internal/repository/memory"
	service "Task-Wallet-Service/internal/service"
)

func main() {
	rep := repository.NewInMemoryUserRepository()
	sv := service.NewUserService(rep)
	sr := server.NewUserHandler(sv)

	sr.StartListen()
}
