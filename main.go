package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/healthl", healthLivenessHandler)
	r.HandleFunc("/healths", healthShallowHandler)
	r.HandleFunc("/healthd", healthDeepHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
