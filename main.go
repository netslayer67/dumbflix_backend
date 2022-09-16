package main

import (
	"fmt"
	"net/http"
	"dumbflix/database"
	"dumbflix/pkg/mysql"
	"dumbflix/routes"

	"github.com/gorilla/mux"
)

func main() {

	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	routes.RouteInit(r.PathPrefix("/api/v1/").Subrouter())

	fmt.Println("Starting API server localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
