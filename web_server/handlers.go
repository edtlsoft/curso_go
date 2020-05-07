package main

import (
	"fmt"
	"net/http"
)

func handleRoot(w http.ResponseWriter, req *htpp.Request){
	fmt.Fprintf(w, "Hello handleRoot")
}