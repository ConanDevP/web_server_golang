package src

import (
	"errors"
	"net/http"
)

type server struct {
	port string
	router *Router
}

func (s *server) Listen() error {
	http.Handle("/",s.router,)
	err := http.ListenAndServe(s.port,s.router)

	if err != nil{
		return errors.New("connection to server failed")
	}
	return nil
}



func NewServer(port string) *server {
	return &server{
		port: port,
		router: NewRouter(),
	}
}
