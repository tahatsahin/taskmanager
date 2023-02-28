package main

import (
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"taskmanager/common"
	"taskmanager/routers"
)

func main() {
	// call startup logic
	common.StartUp()
	// get the mux router
	router := routers.InitRoutes()
	// create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: n,
	}
	log.Println("listening...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("[startServer]: %v", err)
		return
	}
}
