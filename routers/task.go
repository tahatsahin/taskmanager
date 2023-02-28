package routers

import (
	"github.com/gorilla/mux"
	"taskmanager/controllers"
)

func SetTaskRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	// restrict access only to authenticated users
	//router.PathPrefix("/tasks").Handler(negroni.New(
	//	negroni.HandlerFunc(common.Authorize),
	//	negroni.Wrap(taskRouter),
	//))
	return router
}
