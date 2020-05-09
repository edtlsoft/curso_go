package main


import (
	"net/http"
)


type Server struct{
	port string
	router *Router
}


func newServer(port string) *Server{
	return &Server{
		port: port,
		router: newRouter(),
	}
}


func (s *Server) handle(path string, handler http.HandlerFunc){
	s.router.routes[path] = handler
}


func (s *Server) listen() error{
	http.Handle("/", s.router)
	
	error := http.ListenAndServe(s.port, nil)

	return error
}

func (s *Server) addMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc{
	for _, middleware := range middlewares {
		f = middleware(f)
	}

	return f
}