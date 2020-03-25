package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler)

	http.ListenAndServe(":3012", router)
}
