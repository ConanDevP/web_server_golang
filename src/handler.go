package src

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w,"Hello , wordl Handle")
}

func HandlerHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hola, API")
}