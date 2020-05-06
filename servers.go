package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	start := time.Now()

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		checkServer(server)
	}

	elapsedTime := time.Since(start)
	fmt.Printf("Time execution %s\n", elapsedTime)
}

func checkServer(server string){
	_, error := http.Get(server)

	if error != nil {
		fmt.Println(server, "It´s not working")
	}else {
		fmt.Println(server, "It´s working successful")
	}
}