package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	start  := time.Now()
	chanel := make(chan string)

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for {
		if i > 10 {
			break
		}

		for _, server := range servers {
			go checkServer(server, chanel)
		}

		for i:=0; i < len(servers); i++ {
			fmt.Println(<- chanel) 
		}

		time.Sleep(1 * time.Second)

		i++
	}

	elapsedTime := time.Since(start)
	fmt.Printf("Time execution %s\n", elapsedTime)
}

func checkServer(server string, chanel chan string){
	_, error := http.Get(server)

	if error != nil {
		//fmt.Println(server, "It´s not working")
		chanel <- server + " It´s not working"
	}else {
		//fmt.Println(server, "It´s working successful")
		chanel <- server + "It´s working successful"
	}
}