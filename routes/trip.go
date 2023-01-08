package routes

import (
	handlers "dewetour/handler"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gorilla/mux"
)

func TripRoutes(r *mux.Router) {
	tripRepository := repositories.RepositoryTrip(mysql.DB)
	h := handlers.HandlerTrip(tripRepository)

	r.HandleFunc("/trips", h.FindTrip).Methods("GET")
	r.HandleFunc("/trip/{id}", h.GetTrip).Methods("GET")
	r.HandleFunc("/trip", middleware.AuthAdmin(h.CreateTrip)).Methods("POST")
	r.HandleFunc("/trip/{id}", middleware.AuthAdmin(middleware.UploadFile(h.UpdateTrip))).Methods("PATCH")
	r.HandleFunc("/trip/{id}", h.DeleteTrip).Methods("DELETE")

}
