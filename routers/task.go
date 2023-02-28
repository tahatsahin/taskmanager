package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"taskmanager/common"
	"taskmanager/controllers"
)

func SetTaskRoutes(router *mux.Router) *mux.Router {
	taskRouter := mux.NewRouter()
	taskRouter.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")
	// restrict access only to authenticated users
	router.PathPrefix("/tasks").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(taskRouter),
	))
	return router
}
