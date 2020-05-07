package main


import (
	"fmt"
	"net/http"
)


type Router struct{
	rules map[string]http.HandlerFunc
}


func newRouter() *Router{
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}


func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "Hello World!")
}