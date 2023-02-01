package routes

import (
	"housy/handlers"
	"housy/pkg/middleware"
	"housy/pkg/mysql"
	"housy/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", h.FindTransaction).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(middleware.UploadFile(h.CreateTransaction, "attachment"))).Methods("POST")
	// r.HandleFunc("/house/{id}", h.DeleteHouse).Methods("DELETE")
	// r.HandleFunc("/house/{id}", middleware.Auth(middleware.UploadFile(h.UpdateHouse, "image"))).Methods("PATCH")
}
