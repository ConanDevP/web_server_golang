package src

import "net/http"

type Server struct {
	port string
	router *Router
}

/**
=====================Resivers Functions================================
**/

func (s *Server) Listen() {
	http.Handle("/",s.router)
	http.ListenAndServe(s.port, nil)
}

func (s *Server)Handle(path string , handle http.HandlerFunc){
	s.router.rules[path] = handle
}

func (s *Server)AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc{
	for _,m := range middlewares{
		f = m(f)
	}
	return f
}
/**
=====================Functions================================

**/
func NewServer(port string) *Server {
	return &Server{
		port: port,
		router: NewRouter(),
	}
}


