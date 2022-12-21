package routes

import (
	"github.com/PamelaMarinho/go-schedule/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Route) {
	router.HandleFunc("/schedule/", controllers.GetTask).Methods("GET")
	router.HandleFunc("/schedule/{id}", controllers.GetById).Methods("GET")
	router.HandleFunc("/schedule/", controllers.Create).Methods("POST")
	router.HandleFunc("/schedule/{id}", controllers.Update).Methods("PUT")
	router.HandleFunc("/schedule/{id}", controllers.Delete).Methods("DELETE")
}
