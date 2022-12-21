package routes

import (
	"github.com/PamelaMarinho/go-schedule/controllers"
	"github.com/gorilla/mux"
)

var Routes = func(r *mux.Route) {
	r.HandlerFunc("/schedule", controllers.GetTask).Methods("GET")
	r.HandlerFunc("/schedule/{id}", controllers.GetById).Methods("GET")
	r.HandlerFunc("/schedule/{id}", controllers.Update).Methods("PUT")
	r.HandlerFunc("/schedule/{id}", controllers.Delete).Methods("DELETE")
}
