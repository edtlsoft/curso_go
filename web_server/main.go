package main


func main(){
	server := newServer(":3000")

	server.handle("/", handleRoot)
	server.handle("/api", server.addMiddleware(handleHome, CheckAuth()))
	
	server.listen()
}