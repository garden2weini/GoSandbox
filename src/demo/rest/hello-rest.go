package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var tuts = append([]Tut{}, Tut{ID: "1", Isbn: "123", Title: "angular base tut", Author: &Author{
	Name: "zidea",
}})

func main() {
	tuts = append(tuts, Tut{ID: "2", Isbn: "345", Title: "vue base tut", Author: &Author{
		Name: "tina",
	}})
	tuts = append(tuts, Tut{ID: "3", Isbn: "456", Title: "react base tut", Author: &Author{
		Name: "zidea",
	}})

	// Init Router
	r := mux.NewRouter()

	// Sub Router,
	// curl -l http://localhost:4200/rest/tuts
	s := r.PathPrefix("/rest/").Subrouter()
	s.HandleFunc("/tuts", getTuts).Methods("GET")
	// curl -l http://localhost:4200/rest/cat1/001
	s.HandleFunc("/{category}/{id:[0-9]+}", getRegularExpressionParams).Methods("GET")

	// Route Handlers / Endpoints
	// http://localhost:4200/api/tuts
	r.HandleFunc("/api/tuts", getTuts).Methods("GET")
	r.HandleFunc("/api/tuts/{id}", getTut).Methods("GET")
	r.HandleFunc("/api/tuts", createTut).Methods("POST")
	r.HandleFunc("/api/tuts/{id}", updateTut).Methods("PUT")
	r.HandleFunc("/api/tuts/{id}", deleteTut).Methods("DELETE")

	r.Use(loggingMiddleware)

	log.Fatal(http.ListenAndServe(":4200", r))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// tut struct (Model)
type Tut struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Name string `json:"name"`
}

func getRegularExpressionParams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	json.NewEncoder(w).Encode(params)
}

func getTut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	// loop through tuts and find with id
	for _, item := range tuts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Tut{})
}

// get Single Tut
func getTuts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tuts)
}

// create new Tut
func createTut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tut Tut
	_ = json.NewDecoder(r.Body).Decode(&tut)
	tut.ID = strconv.Itoa(rand.Intn(1000)) // Mock Id
	tuts = append(tuts, tut)
	json.NewEncoder(w).Encode(tut)
}

//update tut
func updateTut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call delete handler")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["id"])
	for index, item := range tuts {
		if item.ID == params["id"] {
			tuts = append(tuts[:index], tuts[index+1:]...)
			var tut Tut
			_ = json.NewDecoder(r.Body).Decode(&tut)
			tut.ID = params["id"] // Mock Id
			tuts = append(tuts, tut)
			json.NewEncoder(w).Encode(tut)
			return
		}
	}

	json.NewEncoder(w).Encode(tuts)
}

//delete tut
func deleteTut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call delete handler")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["id"])
	for index, item := range tuts {
		if item.ID == params["id"] {
			tuts = append(tuts[:index], tuts[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(tuts)
}
