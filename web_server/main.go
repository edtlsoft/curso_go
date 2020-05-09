package main


func main(){
	server := newServer(":3000")

	server.handle("/", handleRoot)
	server.handle("/api", handleHome)
	
	server.listen()
}