package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func handleRoot(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Hello handleRoot")
}

func handleHome(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "This is API endpoint")
}

func postRequest(w http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body)

	var metadata Metadata

	error := decoder.Decode(&metadata)

	if error != nil {
		fmt.Fprintf(w, "Error: %v", error)
		return
	}

	fmt.Fprintf(w, "Payload: %v\n", metadata)
}


func userRequest(w http.ResponseWriter, req *http.Request){
	decoder := json.NewDecoder(req.Body)

	var user User

	error := decoder.Decode(&user)

	if error != nil {
		fmt.Fprintf(w, "Error: %v", error)
		return
	}

	response, error := user.toJson()

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(response)
}