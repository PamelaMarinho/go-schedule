package main

import (
	"net/http"

	"github.com/PamelaMarinho/go-schedule/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	http.ListenAndServe("localhost:8080", r)

}
