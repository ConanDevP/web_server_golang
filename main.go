package main


import "github.com/web_server_golang/src"

func main() {
	
	server := src.NewServer(":3000")
	server.Handle("/",src.HandleRoot)
	server.Handle("/api",server.AddMiddleware(src.HandlerHome, src.CheckOut(),src.Logging()))
	server.Listen()


}

