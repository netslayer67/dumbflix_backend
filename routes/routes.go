package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	ProfileRoutes(r)
	categoryRoutes(r)
	AuthRoutes(r)
	FilmRoutes(r)
	TransactionRoutes(r)
}
