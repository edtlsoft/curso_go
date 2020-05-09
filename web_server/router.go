package main


import (
	"net/http"
)


type Router struct{
	routes map[string]map[string]http.HandlerFunc
}


func newRouter() *Router{
	return &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
}


func (r *Router) findHandler(path string, method string) (http.HandlerFunc, bool, bool){
	_, existp := r.routes[path]
	handler, existm := r.routes[path][method]

	return handler, existm, existp
}


func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request){
	handler, existm, existp := r.findHandler(request.URL.Path, request.Method)

	if !existp {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !existm {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}