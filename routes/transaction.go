package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryforTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", middleware.Auth(middleware.IsAdmin(h.FindTransactions))).Methods("GET")
	r.HandleFunc("/transaction/{id}", middleware.Auth(middleware.IsAdmin(h.GetTransaction))).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(middleware.UploadFile(h.CreateTransaction, "attache"))).Methods("POST")
	r.HandleFunc("/transaction/{id}", middleware.Auth(middleware.IsAdmin(h.UpdateTransaction))).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", middleware.Auth(middleware.IsAdmin(h.DeleteTransaction))).Methods("DELETE")

}
