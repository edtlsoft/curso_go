package main


func main(){
	server := newServer(":3000")

	server.handle("GET", "/", handleRoot)
	server.handle("POST", "/api", server.addMiddleware(handleHome, CheckAuth(), Logging()))
	server.handle("POST", "/create", postRequest)
	server.handle("POST", "/user", userRequest)
	
	server.listen()
}