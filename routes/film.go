package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func FilmRoutes(r *mux.Router) {
	filmRepository := repositories.RepositoryFilm(mysql.DB)
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerFilm(filmRepository, categoryRepository)

	r.HandleFunc("/films", h.FindFilm).Methods("GET")
	r.HandleFunc("/film/{id}", h.GetFilm).Methods("GET")
	r.HandleFunc("/film", middleware.Auth(middleware.IsAdmin(middleware.UploadFile(h.CreateFilm, "thumbnailfilm")))).Methods("POST")
	r.HandleFunc("/film/{id}", middleware.Auth(middleware.IsAdmin(h.UpdateFilm))).Methods("PATCH")
	r.HandleFunc("/film/{id}", middleware.Auth(middleware.IsAdmin(h.DeleteFilm))).Methods("DELETE")

}
