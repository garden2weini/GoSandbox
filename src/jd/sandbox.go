package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	//Init Router
	r := mux.NewRouter()
	s := r.PathPrefix("/v1/").Subrouter()
	//Route Handlers / Endpoints
	tryRequest(s)

	//startHttp(r)
	startHttps(r)

}

func startHttp(r *mux.Router) {
	var port = ":9008"
	log.Fatal(http.ListenAndServe(port, r))
}

func startHttps(r *mux.Router) {
	var port = ":10443"
	err := http.ListenAndServeTLS(port, "sandbox.crt", "sandbox.key", r)
	log.Fatal(err)
}
