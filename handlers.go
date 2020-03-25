package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func SuslikHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, <b> "test 1" </b>)
	
}
