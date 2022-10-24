package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func handleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/api/trademarks/word/{name}", test)
	router.HandleFunc("/api/trademarks/word/similar/{name}", test)
	log.Fatal(http.ListenAndServe(":80", router))
}

func main() {
	handleRequests()
}