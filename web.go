package main


import (
	"fmt"
	"net/http"
	"io"
)


type writerWeb struct{}

func (writerWeb) Write(p []byte) (int, error){
	fmt.Println(string(p))

	return len(p), nil
}


func  main(){
	response, error := http.Get("http://google.com")

	if error != nil {
		fmt.Println(error)
	}

	w := writerWeb{}
	io.Copy(w, response.Body)
}