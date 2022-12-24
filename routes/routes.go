package routes

import "github.com/gorilla/mux"

// import (
// 	// "dewetour/handlers"
// 	"dewetour/pkg/mysql"
// 	"dewetour/repositories"
// 	"github.com/gorilla/mux"
//   )

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	TripRoutes(r)
	CountryRoutes(r)
	// Call ProfileRoutes() and ProductRoutes() function here ...
	// ProfileRoutes(r)
	// ProductRoutes(r)
}
