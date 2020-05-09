package main


import (
	"fmt"
	"log"
	"time"
	"net/http"
)


func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request){
			flag := true
			fmt.Println("CheckAtuh")

			if flag {
				f(w, req)
			}else {
				return
			}
		}
	}
}


func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request){
			start := time.Now()
			defer func() {
				log.Println(req.URL.Path, time.Since(start))
			}()

			f(w, req)
		}
	}
}
