package main


import (
	"net/http"
)


type Server struct{
	port string
}


func newServer(port string) *Server{
	return &Server{
		port: port,
	}
}


func (s *Server) listen() error{
	error := http.ListenAndServe(s.port, nil)

	return error
}