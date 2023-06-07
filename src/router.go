package src

import (
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

/**Resiver Functions***/

func (r *Router) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	handler, exist := r.FindHandle(rq.URL.Path)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	handler(w, rq)

}

func (r *Router) FindHandle(path string) (http.HandlerFunc, bool) {
	handle, exist := r.rules[path]

	return handle, exist
}

/**Functions**/
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}
