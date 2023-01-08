package routes

import (
	handlers "dewetour/handler"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gorilla/mux"
)

func CountryRoutes(r *mux.Router) {
	countryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(countryRepository)

	r.HandleFunc("/countries", h.FindCountry).Methods("GET")
	r.HandleFunc("/country/{id}", h.GetCountry).Methods("GET")
	r.HandleFunc("/country", middleware.AuthAdmin(h.CreateCountry)).Methods("POST")
	r.HandleFunc("/country/{id}", middleware.AuthAdmin(h.DeleteCountry)).Methods("DELETE")
	r.HandleFunc("/country/{id}", middleware.AuthAdmin(h.UpdateCountry)).Methods("PATCH")
}
