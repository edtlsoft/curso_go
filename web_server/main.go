package main

import (
	"fmt"
)

func main(){
	server := newServer("3000")
	server.listen()
}