package main


func main(){
	server := newServer(":3000")

	server.handle("GET", "/", handleRoot)
	server.handle("POST", "/api", server.addMiddleware(handleHome, CheckAuth(), Logging()))
	
	server.listen()
}