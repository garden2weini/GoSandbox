package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func tryRequest(s *mux.Router) {
	// http://localhost:9008/jd/helloget?id=maliang
	s.HandleFunc("/jd/helloget", getRequest).Methods("GET")
	// curl localhost:9008/v1/jd/hellopost -X POST -d "param1=comewords&content=articleContent"
	s.HandleFunc("/jd/hellopost", postCommonRequest).Methods("POST")
	// curl localhost:9008/v1/jd/hellojson -X POST -H "Content-Type:application/json" -d '{"param1":"comewords","param2":"articleContent"}'
	s.HandleFunc("/jd/hellojson", postJsonRequest).Methods("POST")
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := r.URL.Query()
	id := vars.Get("id")

	if (len(id)) > 0 {
		json.NewEncoder(w).Encode("NUM:" + id)
		return
	}

	json.NewEncoder(w).Encode("NUM is null")
}

func postCommonRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param1 := r.PostFormValue("param1")

	if (len(param1)) > 0 {
		json.NewEncoder(w).Encode("NUM:" + param1)
		return
	}

	json.NewEncoder(w).Encode("NUM is null")
}

type DemoRequestBody struct {
	Name string `json:"param1"`
	Nike string `json:"param2"`
}

type DemoResponseBody struct {
	Param1 string `json:"arg1"`
	Param2 string `json:"arg2"`
}

func postJsonRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Read failed:", err)
	}
	defer r.Body.Close()

	request := &DemoRequestBody{}
	err = json.Unmarshal(b, request)
	if err != nil {
		log.Println("json format error:", err)
	}

	log.Println("request body:", request)
	response := &DemoResponseBody{}
	response.Param1 = request.Name + "-2"
	response.Param2 = request.Nike + "-2"
	json.NewEncoder(w).Encode(response)

}
