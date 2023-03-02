package routers

import "github.com/gorilla/mux"

// InitRoutes initializes routers
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	// routes for the user
	router = SetUserRoutes(router)
	// routes for the task
	router = SetTaskRoutes(router)
	return router
}
