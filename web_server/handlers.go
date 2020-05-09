package main

import (
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Hello handleRoot")
}

func handleHome(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "This is API endpoint")
}

