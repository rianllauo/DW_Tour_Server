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
	r.HandleFunc("/trip", middleware.UploadFile(h.CreateTrip)).Methods("POST")
	r.HandleFunc("/trip/{id}", h.UpdateTrip).Methods("PATCH")

}
