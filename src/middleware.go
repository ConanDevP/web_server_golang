package src

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func CheckOut() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checkeing Authentication")
			if flag {
				hf(w, r)
				return
			}
			return

		}
	}
}

func Logging() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL, time.Since(start))
			}()
			hf(w, r)
		}

	}
}


