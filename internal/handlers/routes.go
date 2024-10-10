package handlers

import (
	"pleno-go/internal/services"

	"github.com/gorilla/mux"
)

func SetupRoutes(authService *services.AuthService) *mux.Router {
	r := mux.NewRouter()

	authHandler := NewAuthHandler(authService)

	r.HandleFunc("/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	return r
}
